package main

import (
	"fmt"
	"strings"
)

const foo = "foo"

func expand(s string, f func(string) string) string {

	if strings.Contains(s, foo) {
		s = strings.Replace(s, foo, f(foo), -1)
	}
	return s
}

func main() {

	str := "hello"
	str = expand(str, func(s string) string {
		return "yiwei"
	})

	fmt.Println(str)

}
