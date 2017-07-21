package io

import (
	"testing"
	"strings"
	"fmt"
)

func TestUtf8(t *testing.T) {
	//index := utf8Index("a你好a啊", "好")
	index := strings.Index("你好啊", "好")
	fmt.Println(len("你好呀"))
	fmt.Println(index)
}
func utf8Index(str, sub string) (pos int) {
	pos = 0
	asciiPos := strings.Index(str, sub)

	if asciiPos == -1 || asciiPos == 0 {
		pos = asciiPos
		return
	}
	reader := strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		asciiPos -= size
		pos++
		if asciiPos == 0 {
			return
		}
	}
	return
}



