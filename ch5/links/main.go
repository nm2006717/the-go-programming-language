package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := Extract(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2:%v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
//func Extract(url string) ([]string, error) {
//	resp, err := http.Get(url)
//	if err != nil {
//		return nil, err
//	}
//	if resp.StatusCode != http.StatusOK {
//		resp.Body.Close()
//		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
//	}
//	doc, err := html.Parse(resp.Body)
//	resp.Body.Close()
//	if err != nil {
//		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
//	}
//	var links []string
//	visitNode := func(n *html.Node) {
//		if n.Type == html.ElementNode && n.Data == "a" {
//			for _, a := range n.Attr {
//				if a.Key != "href" {
//					continue
//				}
//				link, err := resp.Request.URL.Parse(a.Val)
//				if err != nil {
//					continue // ignore bad URLs
//				}
//				links = append(links, link.String())
//			}
//		}
//	}
//	forEachNode(doc, visitNode, nil)
//	return links, nil
//}

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
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
