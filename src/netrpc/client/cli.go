package main

import (
	"net/rpc"
	"fmt"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:2323")

	if err != nil {
		fmt.Println("connect faild")
		return
	}

	var reply int
	if err = client.Call("Watcher.GetInfo", 1, &reply); err != nil {
		fmt.Println("call faild")
		return
	}
	fmt.Println("result: ", reply)

	if err = client.Call("Add", &Args{23,23}, &reply); err != nil {
		fmt.Println("call faild")
		return
	}
	fmt.Println("result: ", reply)
}

type Args struct {
	a int `json:"user_id" bson:"user_id"`
	b int `json:"user_id" bson:"user_id"`
}