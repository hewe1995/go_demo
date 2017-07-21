//hello.go
package main

import (
	"fmt"
)

//main
func main() {
}
func Test() {
	for i := 0; i < 10; i ++ {
		defer func() {
			fmt.Println("defer execute")
		}()
	}
}
