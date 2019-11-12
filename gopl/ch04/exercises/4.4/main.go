// Copyright © 2019 xue.zeng@outlook.com
// Copyright © 2015 Alan A. A. Donovan · Brian W. Kernighan.
// Published Oct 26, 2015 in paperback and Nov 20 in e-book

// 4.4 编写一个rotate函数，通过一次循环完成旋转
package main

import "fmt"

// rotateLeft rotate slice s left
func rotateLeft(a []int, numRot int) []int {
	lenA := len(a)
	if numRot <= 0 || numRot >= lenA {
		return a
	}
	temp := make([]int, lenA)
	for i, j := 0, numRot; i < lenA; i, j = i+1, j+1 {
		if j == len(a) {
			j = 0
		}
		temp[i] = a[j]
	}
	return temp
}

func rotateRight(s []int, n int) []int {
	n = n % len(s)
	if n > 0 && n < len(s) {
		temp := make([]int, n)
		copy(temp, s[len(s)-n:])

		copy(s[n:], s)
		copy(s, temp)
		return s
	}
	return s
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	// rotateLeft
	rotateLeft(s[:], 2)
	fmt.Println(s) // "[0 1 2 3 4 5]"

	// rotateRight
	rotateRight(s[:], 2)
	fmt.Println(s) // "[4 5 0 1 2 3]"
}
