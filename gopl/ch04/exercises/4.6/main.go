// Copyright © 2019 xue.zeng@outlook.com
// Copyright © 2015 Alan A. A. Donovan · Brian W. Kernighan.
// Published Oct 26, 2015 in paperback and Nov 20 in e-book

// 4.6 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考 unicode.IsSpace）替换成一个空格返回
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// squashSpace squashes each rune of adjacent Unicode spaces in
// UTF-8 encoded []byte slice into a single ASCII space
func squashSpace(bytes []byte) []byte {
	out := bytes[:0]
	var last rune

	for i := 0; i < len(bytes); {
		r, s := utf8.DecodeRune(bytes[i:])

		if !unicode.IsSpace(r) {
			out = append(out, bytes[i:i+s]...)
		} else if unicode.IsSpace(r) && !unicode.IsSpace(last) {
			// ASCII space literal, untyped constant, will transfer
			// to byte when assign to append's parameter
			// ' ' is int32/rune type, character in Go is actually
			// Unicode code point, but string will utf8 encoded
			out = append(out, ' ')
		}
		last = r
		i += s
	}
	return out
}

//!+
func main() {
	a := 'a'
	r := rune(' ')
	fmt.Printf("%T %v %T %v\n", a, a, r, r)

	// string is []byte
	s := "这个\n割裂的\n世界->   hel\n\v\f\tlo \n世界\t\v wo rld \f! \n\v!\n !"
	// output show utf8 encoded
	fmt.Println(len(s), len([]rune(s)))

	b := squashSpace([]byte(s))
	fmt.Println(string(b))
}

//!-
