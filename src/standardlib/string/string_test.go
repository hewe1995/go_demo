package string

import (
	"testing"
	"strings"
	"fmt"
	"strconv"
)

func TestSplit(t *testing.T)  {
	arr := strings.Split("你,好,啊", ",")
	arr1 := strings.SplitAfter("你,好,啊", ",")
	fmt.Println(arr)
	fmt.Println(arr1)
}
func TestByte(t *testing.T)  {
	length := len([]byte("你好呀"))
	fmt.Println(length)
}

func TestStr2Int(t *testing.T)  {
	result, err := strconv.ParseInt("128", 10, 8)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func TestQuote(t *testing.T)  {
	str := strconv.Quote("www.hewe.vip")
	str1 := "www.hewe.vip"
	str2 := `"www.hewe.vip"`
	str3 := "\"www.hewe.vip\""
	fmt.Println(len(str))
	fmt.Println(len(str1))
	fmt.Println(str1)
	fmt.Println(len(str2))
	fmt.Println(str2)
	fmt.Println(len(str3))
	fmt.Println(str3)

}