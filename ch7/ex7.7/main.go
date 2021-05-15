package main

import (
	"./tempconv"
	"flag"
	"fmt"
)

// 练习 7.7：
//	解释为什么帮助信息在它的默认值是20.0没有包含°C的情况下输出了°C。

var temp = tempconv.CelsiusFlag("temp",20.0,"the temperature")

func main() {
	// flag.Parse() 没有解析到任何命令行参数。
	flag.Parse()
	// Celsius实现了String接口，直接打印tmpe里的值。
	fmt.Println(*temp)
}


