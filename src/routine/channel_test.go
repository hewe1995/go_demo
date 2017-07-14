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
		values := <-ch
		fmt.Println(values)
	}
}

func TestChannel2(t *testing.T) {
	num := runtime.NumCPU()
	fmt.Println(num)
}

func reciver(in chan int) {
	fmt.Println(<-in)
}
/**
携程收发双方都在等待对方准备好,如果,先发送数据再开启接受方的携程,会导致死锁.
 */
func TestLock(t *testing.T)  {
	out := make(chan int)
	go reciver(out)
	//这样会导致死锁
	//out <- 2
	//go reciver(out)
	out <- 2
}