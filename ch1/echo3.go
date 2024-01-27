package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// 练习 1.1： 修改 echo 程序，使其能够打印 os.Args[0]，即被执行命令本身的名字。
func test1() {
	var fileNamePath []string = strings.Split(os.Args[0], "\\")
	var fileName string = fileNamePath[len(fileNamePath)-1]
	fmt.Println("FileName:", fileName, "\nArgs:", strings.Join(os.Args[1:], " "))
}

// 练习 1.2： 修改 echo 程序，使其打印每个参数的索引和值，每个一行。
func test2() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println("Index:", i, " Value:", os.Args[i])
	}
}

// 练习 1.2： 修改 echo 程序，使其打印每个参数的索引和值，每个一行。
func test3() {
	var i int = 0
	for _, arg := range os.Args[1:] {
		fmt.Println("Index:", i, " Value:", arg)
		i += 1
	}
}

// **练习 1.3：** 做实验测量潜在低效的版本和使用了 `strings.Join` 的版本的运行时间差异。
func test4() {
	beginTime := time.Now()
	time.Sleep(10)
	var args, sep string
	for i := 1; i < len(os.Args); i++ {
		args += os.Args[i] + sep
		sep = " "
	}
	cost := time.Since(beginTime).Microseconds()
	fmt.Println(args, "Cost:", cost)
}

func main() {
	//test1()
	//test3()
	test4()
}
