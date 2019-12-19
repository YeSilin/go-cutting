package setting

import (
	"encoding/json"
	"fmt"
	"github.com/yesilin/go-cutting/globa"
	"os"
)



// 设置信息编码
//func JsonSettingEncoder(path string, memory, openPs, blackEdge bool, reserve float64) {
func JsonSettingEncoder(path string, newSetting globa.SetInformation) {

	//创建文件（并打开）
	filePtr, err := os.Create(path)
	if err != nil {
		fmt.Println("创建文件失败，err=", err)
		return
	}
	defer filePtr.Close()

	//创建基于文件的JSON编码器
	encoder := json.NewEncoder(filePtr)

	//将新版本信息实例编码到文件中
	err = encoder.Encode(newSetting)
	if err != nil {
		fmt.Println("编码失败，err=", err)

	} else {
		//fmt.Println("编码成功")
	}
}

// 设置信息解码
func JsonSettingDecode(path string) (globa.SetInformation, error) {

	//预定义解码结果
	var oldSetting globa.SetInformation

	//打开文件
	filePtr, _ := os.Open(path)
	defer filePtr.Close()

	//创建该文件的json解码器
	decoder := json.NewDecoder(filePtr)

	//把解码的结果存在oldSetting的地址中
	err := decoder.Decode(&oldSetting)
	if err != nil {
		//fmt.Println("解码失败，err=", err)
		return globa.SetInformation{}, err
	} else {
		//fmt.Printf("解码成功:%#v\n", oldSetting)
		return oldSetting, nil
	}
}
