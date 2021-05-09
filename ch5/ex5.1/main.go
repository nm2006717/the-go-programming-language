package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

//	练习 5.1：
//	修改findlinks代码中遍历n.FirstChild链表的部分，将循环调用visit，改成递归调用。

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	links := visit(doc)
	for _, v := range links {
		fmt.Println(v)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(n *html.Node) (links []string) {
	// 递归出口
	if n == nil {
		return
	}
	links = append(links, visit(n.FirstChild)...)
	links = append(links, visit(n.NextSibling)...)
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	return
}
