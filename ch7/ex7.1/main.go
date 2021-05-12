package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// 	练习 7.1：
//	使用来自ByteCounter的思路，实现一个针对单词和行数的计数器。你会发现bufio.ScanWords非常的有用。

type WordsLineCounter struct {
	lines int
	words int
}

func (c *WordsLineCounter) Write(p []byte) (int, error) {
	reader := bytes.NewReader(p)
	scan := bufio.NewScanner(reader)
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		c.lines++

	}
	reader2 := bytes.NewReader(p)
	scan2 := bufio.NewScanner(reader2)
	scan2.Split(bufio.ScanWords)
	for scan2.Scan() {
		c.words++
	}
	return len(p), nil
}

func main() {
	var w WordsLineCounter
	s := "Hello, World!\nHello, 世界！\nrussia"
	fmt.Fprintf(&w, s)
	fmt.Println(w.words)
	fmt.Println(w.lines)
}
