package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

const (
	EncryptType256 = "sha256"
	EncryptType384 = "sha384"
	EncryptType512 = "sha512"
)

func main() {
	t := flag.String("e", EncryptType256, "选择哈希函数，可选（sha256、sha384、sha512)。默认sha256")
	encryptByte := make([]byte, 0)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() && input.Text() != "q" {
		encryptByte = append(encryptByte, []byte(input.Text())...)
	}

	fmt.Printf("%s\n",encryptByte)
	flag.Parse()

	if t != nil {
		switch *t {
		case EncryptType384:
			fmt.Fprintf(os.Stdout, "%x", sha512.Sum384(encryptByte))
		case EncryptType512:
			fmt.Fprintf(os.Stdout, "%x", sha512.Sum512(encryptByte))
		default:
			fmt.Fprintf(os.Stdout, "%x", sha256.Sum256(encryptByte))
		}
	}
}
