package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

//	练习 5.5：
//	实现countWordsAndImages。（参考练习4.9如何分词）

const (
	image = "img"
	text  = "text"
)

func main() {
	resp, err := http.Get("https://geektutu.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "get geektutu: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "find text element: %v\n", err)
		os.Exit(1)
	}

	maps := make(map[string]int)
	visit(doc, maps)
	fmt.Println(maps)

}

func visit(n *html.Node, maps map[string]int) {
	if n == nil {
		return
	}
	if n.Data != "script" && n.Data != "style" {
		visit(n.FirstChild, maps)
	}
	visit(n.NextSibling, maps)
	if n.Type == html.ElementNode && n.Data == image {
		maps[image] ++
	}
	if n.Type == html.TextNode {
		reader := strings.NewReader(n.Data)
		scan := bufio.NewScanner(reader)
		scan.Split(bufio.ScanWords)
		for scan.Scan() {

			maps[text]++
		}
	}

	return
}
