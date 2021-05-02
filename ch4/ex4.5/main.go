package main

import "fmt"

//	练习 4.5:
//	写一个函数在原地完成消除[]string 中相邻重复的字符串的操作。
//	原地完成消除 / 相邻重复
//	原地消除表示必须在原有的数组上操作
//	遇到相同的先前移一位
//	下标保持不动继续检测当前位置是否跟下一位重复
func main() {
	str := []string{"a", "a", "b", "b", "d", "b", "c", "c"}
	fmt.Println(RemoveDuplicates(str))
}
func RemoveDuplicates(str []string) []string {
	for i := 0; i < len(str)-1; {
		if str[i] == str[i+1] {
			copy(str[i:], str[i+1:])
			str = str[:len(str)-1]
			continue
		}
		i++
	}
	return str
}
