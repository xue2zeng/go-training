// Copyright © 2019 xue.zeng@outlook.com
// Copyright © 2015 Alan A. A. Donovan · Brian W. Kernighan.
// Published Oct 26, 2015 in paperback and Nov 20 in e-book

// 4.8 修改charcount程序，使用unicode.IsLetter等相关的函数，统计字母、数字等 Unicode中不同的字符类别
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	Control = iota
	Letter
	Mark
	Number
	Punct
	Space
	Symbol
	Graphic
	Print
	UTFCatCount
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	var utfcat [UTFCatCount]int     // count of categories of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsPrint(r) {
			utfcat[Print]++
		}

		if unicode.IsGraphic(r) {
			utfcat[Graphic]++
		}

		switch {
		case unicode.IsControl(r):
			utfcat[Control]++
		case unicode.IsLetter(r):
			utfcat[Letter]++
		case unicode.IsMark(r):
			utfcat[Mark]++
		case unicode.IsNumber(r):
			utfcat[Number]++
		case unicode.IsPunct(r):
			utfcat[Punct]++
		case unicode.IsSymbol(r):
			utfcat[Symbol]++
		case unicode.IsSpace(r):
			utfcat[Space]++
		}

		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\nCategory   Count\n")
	// %m.gs, m: width, less<m print spacem >m,then show original char
	// g: control the count of char can be printed
	// -: left align, default: right align
	fmt.Printf("%-7.7s: %4d\n", "Print", utfcat[Print])
	fmt.Printf("%-7.7s: %4d\n", "Graphic", utfcat[Graphic])
	fmt.Printf("%-7.7s: %4d\n", "Control", utfcat[Control])
	fmt.Printf("%-7.7s: %4d\n", "Letter", utfcat[Letter])
	fmt.Printf("%-7.7s: %4d\n", "Mark", utfcat[Mark])
	fmt.Printf("%-7.7s: %4d\n", "Number", utfcat[Number])
	fmt.Printf("%-7.7s: %4d\n", "Punct", utfcat[Punct])
	fmt.Printf("%-7.7s: %4d\n", "Space", utfcat[Space])
	fmt.Printf("%-7.7s: %4d\n", "Symbol", utfcat[Symbol])

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-
