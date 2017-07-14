package main

import (
	"encoding/json"
	"fmt"
)

/*
不是所有的数据都可以编码为 JSON 类型：只有验证通过的数据结构才能被编码：

JSON 对象只支持字符串类型的 key；要编码一个 Go map 类型，map 必须是 map[string]T（T是 json 包中支持的任何类型）
Channel，复杂类型和函数类型不能被编码
不支持循环数据结构；它将引起序列化进入一个无限循环
指针可以被编码，实际上是对指针指向的值进行编码（或者指针是 nil）
 */
import (
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

func go2json() {
	addr := &Address{"company", "beijing", "china"}
	var addres = [5]*Address{&Address{"company", "beijing", "china"}, &Address{"company", "beijing", "china"}}
	js, err := json.Marshal(addres)
	js1, err := json.Marshal(addr)
	if err != nil {
		fmt.Println("err occur: ", err)
	}
	fmt.Println(string(js))
	fmt.Println(string(js1))
}
func go2jsonWithEncoder() {
	addr := &Address{"company", "beijing", "china"}
	var addres = [5]*Address{&Address{"company", "beijing", "china"}, &Address{"company", "beijing", "china"}}
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(addres)
	err = enc.Encode(addr)
	if err != nil {
		fmt.Println("err occur: ", err)
	}
}
func main() {
	//go2json()
	go2jsonWithEncoder()
}
