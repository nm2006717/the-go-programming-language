package main

import (
	"crypto/sha256"
	"fmt"
)

func differBitsCount(sha1 [32]byte, sha2 [32]byte) int {
	cnt := 0
	fmt.Printf("%08b\n", sha1[3])
	a := fmt.Sprintf("%08b\n", sha1)
	b := fmt.Sprintf("%08b\n", sha2)
	fmt.Println(len(a))
	fmt.Println(len(b))
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != b[i] {
			cnt++
		}
	}

	return cnt
}

func main() {
	sha1 := sha256.Sum256([]byte("a"))
	sha2 := sha256.Sum256([]byte("A"))
	difCnt := differBitsCount(sha1, sha2)

	fmt.Printf("different count before sha1 and sha2 is %d\n", difCnt)
}
