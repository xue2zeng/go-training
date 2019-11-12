// Copyright © 2019 xue.zeng@outlook.com
// Copyright © 2015 Alan A. A. Donovan · Brian W. Kernighan.
// Published Oct 26, 2015 in paperback and Nov 20 in e-book

// 4.2 编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag 定制，输出SHA384或SHA512哈希算法。
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

const (
	useSha256 = iota
	useSha384
	useSha512
)

func main() {
	var methodStr string
	mySet := flag.NewFlagSet("", flag.ExitOnError)
	mySet.StringVar(&methodStr, "m", "256", "sha method")
	mySet.Parse(os.Args[1:])
	method := useSha256
	if len(os.Args) > 0 {
		switch methodStr {
		case "384":
			method = useSha384
		case "512":
			method = useSha512
		default:
			method = useSha256
		}
	}

	var input string
	fmt.Println("Enter text:")
	for {
		fmt.Scan(&input)
		if input == "" {
			continue
		}
		value := []byte(input)
		switch method {
		case useSha256:
			hash := sha256.Sum256(value)
			printHash(hash[:])
		case useSha384:
			hash := sha512.Sum384(value)
			printHash(hash[:])
		case useSha512:
			hash := sha512.Sum512(value)
			printHash(hash[:])
		}
	}
}

func printHash(hash []byte) {
	for _, v := range hash {
		fmt.Printf("%X", v)
	}
	fmt.Println()
}
