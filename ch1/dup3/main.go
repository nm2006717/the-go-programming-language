package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 从文件中一次性全部读入内存，然后统计重复行的次数，读到q退出。
func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if line == "q\r" {
				break
			}
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
