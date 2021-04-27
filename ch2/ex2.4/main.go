package main

import (
	"./popcount"
	"fmt"
)

func main() {
	nums := popcount.PopCount(4)
	fmt.Println(nums)
}
