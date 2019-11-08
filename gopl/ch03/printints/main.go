// Copyright © 2019 xue.zeng@outlook.com
// Copyright © 2015 Alan A. A. Donovan · Brian W. Kernighan.
// Published Oct 26, 2015 in paperback and Nov 20 in e-book

// Printints demonstrates the use of bytes.Buffer to format a string.
package main

import (
	"bytes"
	"fmt"
)

//!+
// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3}))
	// Output:
	// 1, 2, 3]
}

//!-
