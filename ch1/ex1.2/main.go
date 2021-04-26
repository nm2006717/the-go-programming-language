package main

import (
	"fmt"
	"os"
)

// 练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。
func main() {
	for k,arg:=range os.Args[1:]{
		fmt.Println(k,arg)
	}
}
