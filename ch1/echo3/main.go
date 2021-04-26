package main

import (
	"fmt"
	"os"
	"strings"
)
// Echo3 prints its command-line arguments.
func main() {
	fmt.Println(strings.Join(os.Args[1:]," "))
	fmt.Println(os.Args[1:])
}
