package main

import (
	"fmt"
)

//	练习 4.5：
//	写一个函数在原地完成消除[]string中相邻重复的字符串的操作。



func remove(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2)) // "[5 6 9 8]
}