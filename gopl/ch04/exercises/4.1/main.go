// Copyright © 2019 xue.zeng@outlook.com
// Copyright © 2015 Alan A. A. Donovan · Brian W. Kernighan.
// Published Oct 26, 2015 in paperback and Nov 20 in e-book

// 4.1 编写一个函数，计算两个SHA256哈希码中不同bit的数目。（参考2.6.2节的 PopCount函数。)
package main

import (
	"crypto/sha256" //!+
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8

	// 125 bits are different between c1 and c2
	fmt.Println(bitsDiff(&c1, &c2))
}

// difference means 'xor'
// only accepted sha256 digest
func bitsDiff(b1, b2 *[sha256.Size]byte) int {
	var sum int
	for i := 0; i < sha256.Size; i++ {
		sum += int(pc[b1[i]^b2[i]])
	}
	return sum
}

//!-
