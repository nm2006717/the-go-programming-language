package main

import (
	"fmt"
	"strings"
)
//	练习5.16：编写多参数版本的strings.Join。sAZ
func join(sep string, args ...string) string {
	if len(args) == 0 {
		return ""
	}
	if len(args) == 1 {
		return args[0]
	}
	var bufStr strings.Builder
	for k, val := range args {
		bufStr.WriteString(val)
		if k != len(args)-1 {
			bufStr.WriteString(sep)
		}
	}
	return bufStr.String()
}

func main() {
	var strs []string
	strs = append(strs, "aaa", "bbb", "ccc", "ddd", "eee")
	fmt.Println(join("-", strs...))
}
