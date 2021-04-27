package main

import "fmt"

const (
	boilingF = 212.0
	symbolF = "℉"
	symbolC = "℃"
)

func main() {
	var f =boilingF
	var c = (f-32) * 5 /9
	fmt.Printf("boiling point = %g%s or %g%s\n",f,symbolF,c,symbolC)
}
