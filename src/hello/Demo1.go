//hello.go
package main

import (
	"os"
	"fmt"
)

//main
func main() {
	file, err := os.OpenFile("ni", 2, 0666)
	if err != nil {

	}
	defer file.Close()
	name := file.Name()
	fmt.Println(name)
}
