package main

import (
	"./popcount"
	"fmt"
)

func main() {
	nums := popcount.PopCount(3)
	fmt.Println(nums)
}