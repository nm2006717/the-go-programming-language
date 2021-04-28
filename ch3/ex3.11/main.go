package main

import (
	"bytes"
	"fmt"
	"strings"
)

//	练习 3.11：
//	完善comma函数，以支持浮点数处理和一个可选的正负号的处理。
func comma(s string) string {
	dotIndex := strings.IndexByte(s, '.')
	if dotIndex != -1 {
		right := s[dotIndex:]
		left := s[:dotIndex]
		return fmt.Sprint(addComma(left), right)
	} else {
		return addComma(s).String()
	}
}

func addComma(left string) *bytes.Buffer {
	var buf bytes.Buffer
	if left[0] == '+' || left[0] == '-' {
		buf.WriteByte(left[0])
		left = left[1:]
	}
	n := len(left)
	if len(left) <= 3 {
		buf.WriteString(left)
		return &buf
	}
	for i := 0; i < n%3; i++ {
		buf.WriteByte(left[i])
	}
	buf.WriteByte(',')
	cnt := 0
	for i := n % 3; i < n; i++ {
		buf.WriteByte(left[i])
		cnt++
		if cnt%3 == 0 && i != n-1 {
			buf.WriteByte(',')
		}
	}
	return &buf
}

func main() {
	oneMillion := "10.555555555555555555555555555557896244"
	fmt.Println(comma(oneMillion))
}
