/*
写一个tar文件,首先定义头部信息和内容,然后先将头信息写入,然后写入对应的内容.
 */
package archive

import (
	"testing"
	"bytes"
	"archive/tar"
	"fmt"
	"os"
	"bufio"
	"io"
)

func TestWrite(t *testing.T) {
	//写文件流
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)

	//定义文件名和内容
	var files = []struct {
		Name, Body string
	}{
		{"README.md", "this is a test tar file"},
		{"ting.txt", "afk;alasfd this is content"},
	}
	//遍历files
	for _, fi := range files {
		hdr := &tar.Header{
			Name: fi.Name,
			Mode: 0600,
			Size: int64(len(fi.Body)),
		}
		//写头
		if err := tw.WriteHeader(hdr); err != nil {
			fmt.Println(err)
		}
		//写内容
		if _, err := tw.Write([]byte(fi.Body)); err != nil {
			fmt.Println(err)
		}
	}
	//关闭tar的写流
	if err := tw.Close(); err != nil {
		fmt.Println(err)
	}
	//打开文件并把buf中的内容写到文件中
	sav, err := os.OpenFile("tee.tar", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer sav.Close()
	sav.Write(buf.Bytes())

}
func TestReader(t *testing.T) {
	//打开文件
	fil, err := os.Open("tee.tar")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fil.Close()

	//缓冲区(可选)
	buf := bufio.NewReader(fil)

	//tar reader
	tr := tar.NewReader(buf)

	for {
		//获取头信息
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("name:	", hdr.Name)
		//获取内容
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println()
	}
}
