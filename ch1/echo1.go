package main

import (
	"fmt"
	"os"
)

func main() {
	// 声明变量s、sep，其中sep为字符串类型
	var s string
	var sep string = " "
	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i] + sep
	}
	fmt.Println(s)
}
