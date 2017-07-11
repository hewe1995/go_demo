package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"io"
	"strconv"
	"time"
	"sortAlg/qsort"
	"sortAlg/bubblesort"
)

var inFile *string = flag.String("i", "", "input file")
var outFile *string = flag.String("o", "outFile", "output file")
var alg *string = flag.String("a", "qsort", "sort alg")

/**
排序主函数
 */
func main() {
	flag.Parse()
	if inFile != nil {
		fmt.Println("inFile: ", *inFile, "outFile: ", *outFile, "alg: ", *alg)
	} else {
		fmt.Println("inFile null !")
		return
	}
	values, err := readFromFile(*inFile)
	if err != nil {
		fmt.Println("error occur: ", err)
		return
	}
	for _, v := range values {
		fmt.Print(v, "	")
	}

	t1 := time.Now()
	switch *alg {
	case "qsort":
		qsort.Qsort(values)
	case "bubblesort":
		bubblesort.BubbleSort(values)
	default:
		fmt.Println("usuport!")
		return

	}
	t2 := time.Now()

	fmt.Println("sort finish , costs: ", t2.Sub(t1))

	err = writeToFile(values, *outFile)
	if err != nil {
		fmt.Println("err occurs: ", err)
	}
}

/**
从文件中读取数据
@parm inFile 文件名
@return values 读取的数据
@return err 出现异常
 */
func readFromFile(inFile string) (values []int, err error) {
	file, err := os.Open(inFile)
	if err != nil {
		fmt.Println("file not found or open exception!")
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				//不是到达文件末尾的异常
				err = err1
			}
			break
		}
		if isPrefix {
			//一行的内容太多
			fmt.Println("line too long")
			return
		}
		//获取数据
		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

/**
写数据到文件中
 */
func writeToFile(values []int, outFile string) (err error) {
	file, err := os.Create(outFile)
	if err != nil {
		fmt.Println("create file faild: ", err)
		return
	}
	defer file.Close()

	for _, v := range values {
		str := strconv.Itoa(v)
		file.WriteString(str + "\n")

	}
	return
}

