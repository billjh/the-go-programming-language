// go test -bench=B -benchmem
// covers Exercise 2.3, 2.4 and 2.5

// BenchmarkExpr-8         2000000000               0.31 ns/op            0 B/op          0 allocs/op
// BenchmarkLoop-8         100000000               17.8 ns/op             0 B/op          0 allocs/op
// BenchmarkShiftBits-8    30000000                60.5 ns/op             0 B/op          0 allocs/op
// BenchmarkClearBits-8    100000000               16.7 ns/op             0 B/op          0 allocs/op

package popcount

import (
	"testing"

	"gopl.io/ch2/popcount"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popCountLoop(x uint64) int {
	count := 0
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(uint(i)*8))])
	}
	return count
}

func popCountShiftBits(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			count++
		}
		x = x >> 1
	}
	return count
}

func popCountClearBits(x uint64) int {
	count := 0
	for x != 0 {
		count++
		x &= (x - 1)
	}
	return count
}

func TestShiftBits(t *testing.T) {
	if popCountShiftBits(0x1111) != 4 {
		t.Error()
	}
}
func TestClearBits(t *testing.T) {
	if popCountClearBits(0x1111) != 4 {
		t.Error()
	}
}

func BenchmarkExpr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCountLoop(0x1234567890ABCDEF)
	}
}

func BenchmarkShiftBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCountShiftBits(0x1234567890ABCDEF)
	}
}

func BenchmarkClearBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCountClearBits(0x1234567890ABCDEF)
	}
}
