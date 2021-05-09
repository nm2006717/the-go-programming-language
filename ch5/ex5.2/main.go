package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

// 	练习 5.2：
//	编写函数，记录在HTML树中出现的同名元素的次数。
func main() {
	resp, err := http.Get("http://geektutu.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "get geektutu: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	counts := make(map[string]int)
	visit(doc, &counts)
	fmt.Println(counts)

}

func visit(n *html.Node, counts *map[string]int) {
	if n == nil {
		return
	}
	visit(n.FirstChild, counts)
	visit(n.NextSibling, counts)
	if n.Type == html.ElementNode {
		(*counts)[n.Data] ++
	}
}
