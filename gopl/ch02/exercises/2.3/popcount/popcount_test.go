package popcount

/*
$ go test -benchmem -run=^$ go-training/gopl/ch02/exercises/2.3/popcount -bench .
goos: darwin
goarch: amd64
pkg: go-training/gopl/ch02/exercises/2.3/popcount
BenchmarkPopCount-8    	1000000000	         0.358 ns/op	       0 B/op	       0 allocs/op
BenchmarkPopCount2-8   	60609717	        19.5 ns/op	       0 B/op	       0 allocs/op
PASS
coverage: 100.0% of statements
ok  	go-training/gopl/ch02/exercises/2.3/popcount	1.606s
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
