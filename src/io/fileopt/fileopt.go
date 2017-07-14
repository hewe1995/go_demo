package main

import (
	"os"
	"fmt"
	"strings"
	"os/exec"
	"bufio"
	"io"
)

/**
垂直读取文件,比如:
hewe 11
ting 23
那么读取的结果是 hewe ting  11 23
 */
func readFileVerticle() {
	file, err := os.Open("./db.txt")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	defer file.Close()

	var col1, col2 []string
	for {
		var v1, v2 string
		_, err := fmt.Fscanln(file, &v1, &v2)
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
	}
	fmt.Println(col1)
	fmt.Println(col2)

}

/**
写内容到文件中,文件不存在则创建文件
 */
func writeOrCreateFile() {
	file, err := os.OpenFile("db.txt", os.O_WRONLY|os.O_CREATE, 0)
	if err != nil {
		fmt.Println("err occur: ", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writerStr := "nihaoya"
	for i := 0; i < 10; i++ {
		writer.WriteString(writerStr)
	}
	writer.Flush()
}

/**
复制文件
 */
func copyFile() {
	file, err := os.Open("db.txt")
	if err != nil {
		fmt.Println("err occur: ", err)
	}
	defer file.Close()
	dst, err := os.OpenFile("db-replication.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("err occur: ", err)
	}
	io.Copy(dst, file)
}
func main() {
	//fmt.Println(getCurrentPath())
	//readFileVerticle()
	//writeOrCreateFile()
	copyFile()
}
func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	checkErr(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0: i+1])
	return path
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
