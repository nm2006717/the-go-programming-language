package main

import "fmt"

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}

	return comma(s[:len(s)-3]) + "," + s[len(s)-3:]
}

func main() {
	oneMillion := "1000000"

	fmt.Println(comma(oneMillion))
}
