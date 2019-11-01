package popcount

/*
$ go test -benchmem -run=^$ go-training/gopl/ch02/exercises/2.4+2.5/popcount -bench .
goos: darwin
goarch: amd64
pkg: go-training/gopl/ch02/exercises/2.4+2.5/popcount
BenchmarkPopCount-8             	1000000000	         0.389 ns/op	       0 B/op	       0 allocs/op
BenchmarkPopCount2-8            	60432021	        20.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkPopCountByShifting-8   	25016593	        47.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkPopCountByClearing-8   	59686516	        19.9 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 100.0% of statements
ok  	go-training/gopl/ch02/exercises/2.4+2.5/popcount	4.123s
Success: Benchmarks passed.
*/

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}
