package main

import (
	"bytes"
	"fmt"
)

//	练习 3.10：
//	编写一个非递归版本的comma函数，
//	使用bytes.Buffer代替字符串链接操作。
func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	for i := 0; i < n%3; i++ {
		buf.WriteByte(s[i])
	}
	buf.WriteByte(',')

	cnt := 0
	for i := n % 3; i < n; i++ {
		buf.WriteByte(s[i])
		cnt++
		if cnt%3 == 0 && i != n-1 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}

func main() {
	oneMillion := "10000000000000000000000"

	fmt.Println(comma(oneMillion))
}
