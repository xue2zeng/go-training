// Copyright © 2019 xue.zeng@outlook.com
// Copyright © 2015 Alan A. A. Donovan · Brian W. Kernighan.
// Published Oct 26, 2015 in paperback and Nov 20 in e-book

// 4.9 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用 Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // counts of Unicode characters
	input := bufio.NewScanner(os.Stdin)

	// set SplitFunc with ScanWords instead of default ScanLines
	// ScanWords just return space-separated words
	input.Split(bufio.ScanWords)

	for input.Scan() {
		counts[input.Text()]++
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "WordCount: %v", err)
		os.Exit(1)
	}

	for k, v := range counts {
		fmt.Println(k, v)
	}
}

//!-
