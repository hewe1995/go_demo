package main

import "time"

/**
读取channel的超时处理
 */
func main_test() {
	timeout := make(chan bool, 1)
	ch := make(chan int, 1) //数据channel
	go func() {
		time.Sleep(1e9) // 1的9次方, 这里的单位是纳秒
		timeout <- true
	}()
	select {
	case <-ch:
		//获取到数据
	case <-timeout:
		//超时
	}
}
