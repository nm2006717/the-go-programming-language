package main

import (
	"flag"
	"fmt"
	"strings"
)

//	n:忽略尾随的换行符开关
var n = flag.Bool("n", false, "omit trailing newline")

// sep 命令行参数之间的连接符
var sep = flag.String("s", " ", "separator")

//	当程序运行时，必须在使用标志参数对应的变量之前先调用flag.Parse函数，用于更新每个标志参数对应变量的值（之前是默认值）。
//	对于非标志参数的普通命令行参数可以通过调用flag.Args()函数来访问，返回值对应一个字符串类型的slice。
//	如果在flag.Parse函数解析命令行参数时遇到错误，默认将打印相关的提示信息，然后调用os.Exit(2)终止程序。
func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
