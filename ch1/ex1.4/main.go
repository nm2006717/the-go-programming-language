package main

import (
	"bufio"
	"fmt"
	"os"
)

// 修改dup2，出现重复的行时打印文件名称。
func main() {
	countMaps := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, countMaps)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, countMaps)
			f.Close()
		}
	}
	for name, counts := range countMaps {
		for _, n := range counts {
			if n > 1 {
				fmt.Printf("文件%s中有重复行\n", name)
				break
			}
		}
	}
}

func countLines(f *os.File, countMaps map[string]map[string]int) {
	counts := make(map[string]int)
	countMaps[f.Name()] = counts
	input := bufio.NewScanner(f)
	for input.Scan() && input.Text() != "q" {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
