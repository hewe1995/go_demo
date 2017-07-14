/**
从命令行读取数据
 */
package main

import (
	"os"
	"strings"
	"fmt"
	"flag"
	"bufio"
)

/*
获取命令行的参数
 */
func cmd_args() {
	who := "Alice"
	//fmt.Println(os.Args[0])
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], "  ")
	}
	fmt.Println("nihao: ", who)
}

/*
使用flag 获取命令行参数
 */
var newLine = flag.Bool("n", false, "print new line")

const space = " "
const enter = "\n"

/*
使用flag 获取命令行参数
 */
func flag_args() {
	flag.PrintDefaults()
	flag.Parse()
	var s string = ""
	fmt.Println("args number: ", flag.NArg())
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *newLine {
				s += enter
			}
		}
		s += flag.Arg(i)
	}

	fmt.Println(s)

}

/*
用 buffer 读取文件
 */
func bufReader(r *bufio.Reader) {
	for {
		//读到 \n 停止
		buf, err := r.ReadString('\n')
		if err != nil {
			//fmt.Fprintf(os.Stdout, "%s", buf)
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
	return
}
func bufRead() {
	flag.PrintDefaults()
	flag.Parse()
	if flag.NArg() == 0 {
		bufReader(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		bufReader(bufio.NewReader(f))
	}
}
func main() {
	//cmd_args()
	//flag_args()
	//bufRead()
	sliceRead()
}

func sliceRead() {
	flag.Parse() // Scans the arg list and sets up flags
	if flag.NArg() == 0 {
		sliceReader(os.Stdin)
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if f == nil {
			fmt.Fprintf(os.Stderr, "cat: can't open %s: error %s\n", flag.Arg(i), err)
			os.Exit(1)
		}
		sliceReader(f)
		f.Close()
	}
}
func sliceReader(f *os.File) {
	const NBUF = 512
	var buf [NBUF]byte
	for {
		switch nr, err := f.Read(buf[:]); true { //可以没有true,但是不能没有分号,不然会以为在判断前面的条件
		case nr < 0:
			fmt.Println("error occur: ", err)
		case nr == 0:
			//EOF
			return
		case nr > 0:
			if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing: %s\n", ew.Error())
			}

		}
	}
}
