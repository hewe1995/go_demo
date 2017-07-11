package main

import (
	"net/http"
	"net/rpc"
	"net"
	"fmt"
)

type Watcher int

func (m *Watcher) GetInfo(arg int, result *int) (err error) {
	*result = 1
	return nil
}
func Add(a, b int,sum *int) (err error) {
	*sum = a + b
	return nil
}

func main() {
	watcher := new(Watcher)
	rpc.Register(watcher)
	rpc.Register(Add)
	rpc.HandleHTTP()

	listen, error := net.Listen("tcp", ":2323")
	if error != nil {
		fmt.Println("listen faild")
		return
	}
	fmt.Println("listening ...")
	http.Serve(listen, nil)

}