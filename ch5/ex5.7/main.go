package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

var depth int

func main() {
	//resp, err := http.Get("https://geektutu.com")
	//if err != nil {
	//	log.Fatalf("access get request err: %v\n", err)
	//	return
	//}
	//defer resp.Body.Close()

	file, _ := os.Open("test.html")

	doc, err := html.Parse(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}

// forEachNode针对每个结点x,都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前,pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		if len(n.Attr) == 0 {
			fmt.Printf("%*s<%s", depth*2, "", n.Data)
		} else {
			fmt.Printf("%*s<%s ", depth*2, "", n.Data)
			for i := 0; i < len(n.Attr); i++ {
				fmt.Printf("%s='%s'", n.Attr[i].Key, n.Attr[i].Val)
				if i != len(n.Attr)-1 {
					fmt.Printf(" ")
				}
			}
		}
		if n.FirstChild != nil && n.FirstChild.Type == html.TextNode && n.FirstChild.Data != "" {
			fmt.Printf(">\n")
		} else {
			fmt.Printf("/>\n")
		}
		depth++
	case html.TextNode:
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	case html.CommentNode:
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil && n.FirstChild.Type == html.TextNode && n.FirstChild.Data != "" {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
