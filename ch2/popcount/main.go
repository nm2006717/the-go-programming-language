package main

import (
	"./popcount"
	"fmt"
)

func main() {
	nums := popcount.PopCount(50)
	fmt.Println(nums)
}
