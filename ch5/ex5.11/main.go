package main

import (
	"fmt"
	"log"
	"sort"
)

//var prereqs = map[string][]string{
//	"algorithms": {"data structures"},
//	"calculus":   {"linear algebra"},
//
//	"compilers": {
//		"data structures",
//		"formal languages",
//		"computer organization",
//	},
//
//	"data structures":       {"discrete math"},
//	"database":              {"data structures"},
//	"discrete math":         {"intro to programming"},
//	"formal languages":      {"discrete math"},
//	"networks":              {"operating systems"},
//	"operating systems":     {"data structures", "computer organization"},
//	"programming languages": {"data structures", "computer organization"},
//
//	"linear algebra": {"calculus"},
//}

var prereqs2 = map[string][]string{
	"算法":   {"数据结构"},
	"微积分":  {"线性代数"},
	"线性代数": {"微积分"},
	"编译原理": {
		"数据结构",
		"形式语言 ",
		"计算机体系结构",
	},
	"数据结构":   {"离散数学"},
	"数据库":    {"数据结构"},
	"离散数学":   {"编程导论"},
	"形式语言":   {"离散数学"},
	"网络":     {"操作系统"},
	"操作系统":   {"数据结构", "计算机体系结构"},
	"程序设计语言": {"数据结构", "计算机体系结构"},
}

func main() {
	sorted, err := topoSort(prereqs2)
	if err != nil {
		log.Println(err)
	}
	for i, course := range sorted {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string) error

	visitAll = func(items []string) error {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				if err := visitAll(m[item]); err != nil {
					return err
				}
				order = append(order, item)
			} else {
				hasCycle := true
				for _, s := range order {
					if s == item {
						hasCycle = false
					}
				}
				if hasCycle {
					return fmt.Errorf("has cycle: %s", item)
				}
			}
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	if err := visitAll(keys); err != nil {
		return nil, err
	}

	return order, nil
}
