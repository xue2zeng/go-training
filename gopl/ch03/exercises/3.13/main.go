// Copyright © 2019 xue.zeng@outlook.com
// Copyright © 2015 Alan A. A. Donovan · Brian W. Kernighan.
// Published Oct 26, 2015 in paperback and Nov 20 in e-book

//!+
// 3.13 编写KB、MB的常量声明，然后扩展到YB
package main

import "fmt"

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB
	GiB
	TiB // overflow 32 int
	PiB
	EiB
	ZiB // overflow 64 int
	YiB
)

// seems the only valid way
const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
)

func main() {
	fmt.Printf("%T %[1]v\n", KiB)
	fmt.Printf("%T %[1]v\n", MiB)
	fmt.Printf("%T %[1]v\n", GiB)
	fmt.Printf("%T %[1]v\n", TiB)
	fmt.Printf("%T %[1]v\n", PiB)
	fmt.Printf("%T %[1]v\n", EiB)
	//fmt.Printf("%T %[1]v\n", ZiB)

	fmt.Printf("%T %[1]v\n", KB)
	fmt.Printf("%T %[1]v\n", MB)
	fmt.Printf("%T %[1]v\n", GB)
	fmt.Printf("%T %[1]v\n", TB)
	fmt.Printf("%T %[1]v\n", PB)
	fmt.Printf("%T %[1]v\n", EB)
	//fmt.Printf("%T %[1]v\n", ZB)

	// impressive
	fmt.Println(YiB / ZiB)
	fmt.Println(YB / ZB)

	//
	var f float64 = 1 + 0i
	f = 'a'
	fmt.Println(f)
}

//!-
