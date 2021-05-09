package main

import "fmt"

//	练习5.15：
//	编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。

func max(vals ...int) int {
	max := 0
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}
func min(vals ...int) int {
	min := 0
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func min1(val int, vals ...int) int {
	for _, v := range vals {
		if v < val {
			val = v
		}
	}
	return val
}

func max1(val int, vals ...int) int {
	for _, v := range vals {
		if v > val {
			val = v
		}
	}
	return val
}

func main() {
	fmt.Println(min())
	fmt.Println(max())

	fmt.Println(min1(1, 2, 3, 4, 5))

	fmt.Println(max1(1, 2, 3, 4, 5))
}
