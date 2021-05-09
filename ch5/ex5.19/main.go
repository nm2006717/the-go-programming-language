
//	练习5.19：
//	使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。

package main

import (
"fmt"
)

func main() {
	fmt.Println(returnNoZero())
}

// 原理：recover 后不会继续执行，而是直接调用 return
func returnNoZero() (result int) {
	defer func() {
		_ = recover()
	}()
	result = 3
	panic("panic!")
}