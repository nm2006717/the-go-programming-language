package main

import (
	"bufio"
	"fmt"
	"os"
)

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
// 从标准输入中统计重复行的数量，遇到字符串"q"退出。

// Printf 参数verb
//%d          十进制整数
//%x, %o, %b  十六进制，八进制，二进制整数。
//%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
//%t          布尔：true或false
//%c          字符（rune） (Unicode码点)
//%s          字符串
//%q          带双引号的字符串"abc"或带单引号的字符'c'
//%v          变量的自然形式（natural format）
//%T          变量的类型
//	%%          字面上的百分号标志（无操作数）
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() && input.Text() != "q" {
		counts[input.Text()]++
	}

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
