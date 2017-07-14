package main

/**
用户输入测试demo
 */
import (
	"fmt"
	"bufio"
	"os"
)

var (
	fname, lastname string
	a               float64
	b               int
	c               string
)
/**
fmt.scan 的使用
 */
func scan() {
	fmt.Println("input_Name:")
	fmt.Scanln(&fname, &lastname)
	fmt.Println(fname, lastname)
	_, err := fmt.Scanf("%f/%d/%s", &a, &b, &c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a, b, c)
}

func main() {
	//scan()
	bufReader()
}

/**
缓冲读取
 */
func bufReader() {
	fmt.Println("inputRemark: ")
	inputReader := bufio.NewReader(os.Stdin)

	input, err := inputReader.ReadString('!')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(input)
	//字符串以"\n"结尾
	if input == "asdf\n" {
		fmt.Println(input + "nihao")
	}
	//


}
