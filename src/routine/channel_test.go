package main

import (
	"testing"
	"fmt"
	"runtime"
)

func TestChannel1(t *testing.T) {
	ch := make(chan int, 1)

	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}
		values := <- ch
		fmt.Println(values)
	}
}

func TestChannel2(t *testing.T)  {
	num := runtime.NumCPU()
	fmt.Println(num)
}