// Copyright © 2019 xue.zeng@outlook.com
// Copyright © 2015 Alan A. A. Donovan · Brian W. Kernighan.
// Published Oct 26, 2015 in paperback and Nov 20 in e-book

//!+

// Comma prints its argument numbers with a comma at each power of 1000.
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
// 3.11 完善comma函数，以支持浮点数处理和一个可选的正负号的处理。
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var sign string

	// skip optional sign flag
	if hasSign := strings.HasPrefix(s, "+") ||
		strings.HasPrefix(s, "-"); hasSign {
		sign = s[0:1]
		s = s[1:]
	}

	var buf bytes.Buffer
	buf.WriteString(sign)

	if period := strings.LastIndex(s, "."); period != -1 {
		doWork(s[:period], &buf)
		buf.WriteByte('.')
		doWork(s[period+1:], &buf)
	} else {
		doWork(s, &buf)
	}

	return buf.String()
}

// add to only one buf
func doWork(s string, buf *bytes.Buffer) {
	n := len(s)
	if n <= 3 {
		buf.WriteString(s)
		return
	}

	d, t := len(s)%3, len(s)/3
	if d != 0 {
		buf.WriteString(s[:d])
		buf.WriteByte(',')
	}

	for i := 0; i < t; i++ {
		buf.WriteString(s[i*3+d : d+i*3+3])
		if i != t-1 {
			buf.WriteByte(',')
		}
	}
}

//!-
