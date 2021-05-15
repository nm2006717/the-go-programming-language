package main

import (
	"./tempconv"
	"flag"
	"fmt"
)

//	练习 7.6：
//	对tempFlag加入支持开尔文温度。

var temp = tempconv.CelsiusFlag("temp",20.0,"the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}


