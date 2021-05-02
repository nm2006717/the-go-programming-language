package main

import (
	"fmt"
	"unicode"
)

//	练习 4.6:
//	编写一个函数,原地将一个 UTF-8 编码的[]byte 类型的 slice 中相邻的空格(参考 unicode.IsSpace)替换成一个空格返回
func main() {
	bs := []byte("abc     a aaa     ccc  ddd d")
	s := RemoveSpace(bs)
	fmt.Println(string(s))

}
func RemoveSpace(s []byte) []byte{
	for i := 0; i < len(s)-1; {
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[i+1])) {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			continue
		}
		i++
	}
	return s
}
