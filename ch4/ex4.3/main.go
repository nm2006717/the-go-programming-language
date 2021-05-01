package main

import "fmt"

// 练习 4.3：
// 重写reverse函数，使用数组指针代替slice。
func reverse(nums *[]int)   {
	for i, j := 0, len(*nums)-1; i < j; i, j = i+1, j-1 {
		(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
	}
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"
}
