package main

import (
	"fmt"
	"time"
)

//	练习 3.12：
//	编写一个函数，判断两个字符串是否是相互打乱的，也就是说它们有着相同的字符，
//	但是对应不同的顺序。

func isDisorderString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	maps := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		maps[s1[i]]++
	}

	for i := 0; i < len(s2); i++ {
		if _, ok := maps[s2[i]]; ok {
			maps[s2[i]]--
		}
	}
	for _, v := range maps {
		if v != 0 {
			return false
		}
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isDisorderString("abc","bac"))

	const timeout = 5 * time.Minute
	const noDelay time.Duration = 0


}
