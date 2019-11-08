// Copyright © 2019 xue.zeng@outlook.com
// Copyright © 2015 Alan A. A. Donovan · Brian W. Kernighan.
// Published Oct 26, 2015 in paperback and Nov 20 in e-book

// Comma prints its argument numbers with a comma at each power of 1000.
package main

import (
	"fmt"
	"os"
)

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

//!-

//
// Example:
// 	$ go build go-training/gopl/ch03/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf(" %s\n", comma(os.Args[i]))
	}
}
