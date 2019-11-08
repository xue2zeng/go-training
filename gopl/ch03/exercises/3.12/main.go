// Copyright © 2019 xue.zeng@outlook.com
// Copyright © 2015 Alan A. A. Donovan · Brian W. Kernighan.
// Published Oct 26, 2015 in paperback and Nov 20 in e-book

//!+
// 3.12 编写一个函数，判断两个字符串是否是是相互打乱的，也就是说它们有着相同的 字符，但是对应不同的顺序。
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please input 2 strings")
		os.Exit(1)
	}

	s1 := os.Args[1]
	s2 := os.Args[2]
	if areAnagrams(s1, s2) {
		fmt.Println("They are anagrams")
	} else {
		fmt.Println("They are not anagrams")
	}
}

func areAnagrams(s1, s2 string) bool {
	if s1 == s2 || len(s1) != len(s2) {
		return false
	}

	s1 = SortString(s1)
	s2 = SortString(s2)
	if s1 == s2 {
		return true
	}
	return false
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
