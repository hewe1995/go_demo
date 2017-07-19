//hello.go
package main

import (
	"fmt"
	"path"
)

//main
func main() {
	path := path.Join("parent", "child")
	fmt.Println(path)
}
func Test() {
	for i := 0; i < 10; i ++ {
		defer func() {
			fmt.Println("defer execute")
		}()
	}
}
