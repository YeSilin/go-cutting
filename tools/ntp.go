
package tools

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
)

const (
	UNIX_STA_TIMESTAMP = 2208988800
)

/**
NTP协议 http://www.ntp.org/documentation.html
@author mengdj@outlook.com
*/
type ntp struct {
	//1:32bits
	Li        uint8 //2 bits
	Vn        uint8 //3 bits
	Mode      uint8 //3 bits
	Stratum   uint8
	Poll      uint8
	Precision uint8
	//2:
	RootDelay           int32
	RootDispersion      int32
	ReferenceIdentifier int32
	//64位时间戳
	ReferenceTimestamp uint64 //指示系统时钟最后一次校准的时间
	OriginateTimestamp uint64 //指示客户向服务器发起请求的时间
	ReceiveTimestamp   uint64 //指服务器收到客户请求的时间
	TransmitTimestamp  uint64 //指示服务器向客户发时间戳的时间
}

func newNtp() (p *ntp) {
	//其他参数通常都是服务器返回的
	p = &ntp{Li: 0, Vn: 3, Mode: 3, Stratum: 0}
	return p
}

/**
构建NTP协议信息
*/
func (this *ntp) GetBytes() []byte {
	//注意网络上使用的是大端字节排序
	buf := &bytes.Buffer{}
	head := (this.Li << 6) | (this.Vn << 3) | ((this.Mode << 5) >> 5)
	binary.Write(buf, binary.BigEndian, uint8(head))
	binary.Write(buf, binary.BigEndian, this.Stratum)
	binary.Write(buf, binary.BigEndian, this.Poll)
	binary.Write(buf, binary.BigEndian, this.Precision)
	//写入其他字节数据
	binary.Write(buf, binary.BigEndian, this.RootDelay)
	binary.Write(buf, binary.BigEndian, this.RootDispersion)
	binary.Write(buf, binary.BigEndian, this.ReferenceIdentifier)
	binary.Write(buf, binary.BigEndian, this.ReferenceTimestamp)
	binary.Write(buf, binary.BigEndian, this.OriginateTimestamp)
	binary.Write(buf, binary.BigEndian, this.ReceiveTimestamp)
	binary.Write(buf, binary.BigEndian, this.TransmitTimestamp)
	//[27 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	return buf.Bytes()
}

func (this *ntp) Parse(bf []byte, useUnixSec bool) {
	var (
		bit8  uint8
		bit32 int32
		bit64 uint64
		rb    *bytes.Reader
	)
	//貌似这binary.Read只能顺序读，不能跳着读，想要跳着读只能使用切片bf
	rb = bytes.NewReader(bf)
	binary.Read(rb, binary.BigEndian, &bit8)
	//向右偏移6位得到前两位LI即可
	this.Li = bit8 >> 6
	//向右偏移2位,向右偏移5位,得到前中间3位
	this.Vn = (bit8 << 2) >> 5
	//向左偏移5位，然后右偏移5位得到最后3位
	this.Mode = (bit8 << 5) >> 5
	binary.Read(rb, binary.BigEndian, &bit8)
	this.Stratum = bit8
	binary.Read(rb, binary.BigEndian, &bit8)
	this.Poll = bit8
	binary.Read(rb, binary.BigEndian, &bit8)
	this.Precision = bit8

	//32bits
	binary.Read(rb, binary.BigEndian, &bit32)
	this.RootDelay = bit32
	binary.Read(rb, binary.BigEndian, &bit32)
	this.RootDispersion = bit32
	binary.Read(rb, binary.BigEndian, &bit32)
	this.ReferenceIdentifier = bit32

	//以下几个字段都是64位时间戳(NTP都是64位的时间戳)
	binary.Read(rb, binary.BigEndian, &bit64)
	this.ReferenceTimestamp = bit64
	binary.Read(rb, binary.BigEndian, &bit64)
	this.OriginateTimestamp = bit64
	binary.Read(rb, binary.BigEndian, &bit64)
	this.ReceiveTimestamp = bit64
	binary.Read(rb, binary.BigEndian, &bit64)
	this.TransmitTimestamp = bit64
	//转换为unix时间戳,先左偏移32位拿到64位时间戳的整数部分，然后ntp的起始时间戳 1900年1月1日 0时0分0秒 2208988800
	if useUnixSec {
		this.ReferenceTimestamp = (this.ReceiveTimestamp >> 32) - UNIX_STA_TIMESTAMP
		if this.OriginateTimestamp > 0 {
			this.OriginateTimestamp = (this.OriginateTimestamp >> 32) - UNIX_STA_TIMESTAMP
		}
		this.ReceiveTimestamp = (this.ReceiveTimestamp >> 32) - UNIX_STA_TIMESTAMP
		this.TransmitTimestamp = (this.TransmitTimestamp >> 32) - UNIX_STA_TIMESTAMP
	}
}

// 获取NTP时间戳
func GetNtpTime() int64 {

	var (
		ntp    *ntp
		buffer []byte
		err    error
		ret    int
	)
	//链接阿里云NTP服务器,NTP有很多免费服务器可以使用time.windows.com
	conn, err := net.Dial("udp", "ntp1.aliyun.com:123")
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
		conn.Close()
	}()
	ntp = newNtp()
	conn.Write(ntp.GetBytes())
	buffer = make([]byte, 2048)
	ret, err = conn.Read(buffer)


	if err == nil {
		if ret > 0 {
			ntp.Parse(buffer, true)
			//fmt.Println(fmt.Sprintf(
			//	"LI:%d\r\n版本:%d\r\n模式:%d\r\n精度:%d\r\n轮询:%d\r\n系统精度:%d\r\n延时:%ds\r\n最大误差:%d\r\n时钟表示:%d\r\n时间戳:%d %d %d %d\r\n",
			//	ntp.Li,
			//	ntp.Vn,
			//	ntp.Mode,
			//	ntp.Stratum,
			//	ntp.Poll,
			//	ntp.Precision,           // 系统精度
			//	ntp.RootDelay,           // 延时
			//	ntp.RootDispersion,      // 最大误差
			//	ntp.ReferenceIdentifier, // 时钟表示
			//	ntp.ReferenceTimestamp,  // 参考时间戳
			//	ntp.OriginateTimestamp,  // 起始时间戳
			//	ntp.ReceiveTimestamp,    // 接收时间戳
			//	ntp.TransmitTimestamp,   // 传输时间戳
			//))

			// 返回参考时间,系统时钟最后一次校准的时间
			return int64(ntp.ReferenceTimestamp)
		}
	}
	// 起始时间戳 1900年1月1日 0时0分0秒 2208988800
	return 2208988800  // 获取失败就返回默认时间
}

