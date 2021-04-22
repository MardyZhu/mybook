package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// 定义的和json对接的结构体，首字母必须大写
type Result struct {
	Customer *CustomerI `json:"customer"`
}
type CustomerI struct {
	CustomerId int64    `json:"customerId"`
	Catalog    string   `json:"catalog"`
	Nicknames  []string `json:"nicknames"`
}

func main() {
	filePtr, err := os.Open("./f.json")
	if err != nil {
		fmt.Println("文件打开失败 [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()
	byteValue, _ := ioutil.ReadAll(filePtr)

	var resp Result
	err = json.Unmarshal([]byte(byteValue), &resp)
	// 创建json解码器
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")
	}

	// 以json格式展示出来
	ret, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println(string(ret))
}
