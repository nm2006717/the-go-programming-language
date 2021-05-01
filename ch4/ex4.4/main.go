package main

import "fmt"

//	练习 4.4：
//	编写一个rotate函数，通过一次循环完成旋转。

// 思路:
//	1、先copy数组到tmpNums
//	2、两个指针left,right，left一个指向元素原位置,right指向旋转后钙元素的位置。
//  3、从tmpNums取值，将原位置上的值，存放于旋转后的值。
//  4、left、right指针移动。
//  5、重复3、4步操作，直至left >= cnt
func rotate(nums []int, k int) {
	cnt := len(nums)
	left := 0
	right := (left + k) % cnt
	tmpNums := make([]int, cnt)
	copy(tmpNums, nums)
	for left < cnt {
		nums[right] = tmpNums[left]
		left++
		right = (left + k) % cnt
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}
