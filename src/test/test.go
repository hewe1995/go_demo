package test

import "fmt"

var count int

func init() {
	fmt.Println("init...")
}

func Add()  {
	count = count + 1
}

func Count()  {
	fmt.Println(count)
}