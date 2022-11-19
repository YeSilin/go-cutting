package tools

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Note 笔记结构体
type Note struct {
	Tl      string `json:"tl"`
	Content string `json:"content"`
}

// NoteContent 自动生成的 https://mholt.github.io/json-to-go/
type NoteContent struct {
	Num5 []struct {
		Num3 string `json:"3"`
		Num5 []struct {
			Num2 string `json:"2"`
			Num3 string `json:"3"`
			Num7 []struct {
				Num8 string `json:"8"`
			} `json:"7"`
		} `json:"5"`
	} `json:"5"`
}

// GetYNote 获取有道云笔记
func GetYNote(id string) string {
	// 生成client 参数为默认
	client := &http.Client{}

	// 生成要访问的url
	url := "https://note.youdao.com/yws/api/note/" + id + "?sev=j1&editorType=1&editorVersion=new-json-editor"

	// 提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// 增加请求头部
	reqest.Header.Add("Accept", "*/*")
	reqest.Header.Add("X-Requested-With", "XMLHttpRequest")
	reqest.Header.Add("Referer", "https://note.youdao.com/ynoteshare/index.html?id="+id)
	reqest.Header.Add("Accept-Language", "zh-CN")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko")
	reqest.Header.Add("Host", "note.youdao.com")
	reqest.Header.Add("Connection", "Keep-Alive")
	reqest.Header.Add("Cache-Control", "no-cache")

	// 处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	//// 读取全部
	//respByte, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 打印结果
	//fmt.Println(string(respByte))

	// 用来接收解码后的结果
	var note Note
	err = json.NewDecoder(response.Body).Decode(&note)
	if err != nil {
		panic(err)
	}

	// 对内容进行二次解码
	var noteContent NoteContent
	err = json.Unmarshal([]byte(note.Content), &noteContent)
	if err != nil {
		panic(err)
	}

	// 取出层层叠叠的数据
	var buffer bytes.Buffer
	for i := 0; i < len(noteContent.Num5); i++ {
		//fmt.Println(noteContent.Num5[i].Num5[0].Num7)
		for j := 0; j < len(noteContent.Num5[i].Num5[0].Num7); j++ {
			//fmt.Print(noteContent.Num5[i].Num5[0].Num7[j].Num8)
			buffer.WriteString(noteContent.Num5[i].Num5[0].Num7[j].Num8)
		}
		// 最后一行不默认加换行
		if i != len(noteContent.Num5)-1 {
			//fmt.Println()
			buffer.WriteString("\n")
		}
	}

	// 返回数据
	return buffer.String()
}
