package main

import (
	"fmt"
	"os"
)
// Echo2 prints its command-line arguments.
func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
