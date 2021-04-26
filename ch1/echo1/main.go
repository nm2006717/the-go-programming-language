package main

import (
	"fmt"
	"os"
)
// Echo1 prints its command-line arguments.
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
