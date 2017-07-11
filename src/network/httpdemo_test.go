package network

import (
	"testing"
	"net/http"
	"fmt"
	"os"
	"io"
)

var g = 23

func TestHttp(t *testing.T) {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("err occurs: ", err)
		return
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
func TestMark(t *testing.T) {
	var arr [3]int
	arr[2] = 23
	var arr2 []int
	if arr2 == nil {
		fmt.Println("nil")
	}
	arr2 = arr

}
