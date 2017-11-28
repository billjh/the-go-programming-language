// go test -bench=Echo -benchmem

package ch1

import (
	"strings"
	"testing"
)

type echoFunc func([]string) string

func echo1(args []string) string {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

func echo3(args []string) string {
	return strings.Join(args, " ")
}

func genArgs(n int) []string {
	args := make([]string, n)
	for i := 0; i < n; i++ {
		args[i] = "someArg"
	}
	return args
}

func bench(n int, fn echoFunc) func(*testing.B) {
	return func(b *testing.B) {
		args := genArgs(n)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			fn(args)
		}
	}
}

func BenchmarkEcho(b *testing.B) {
	b.Run("echo1,args=10", bench(10, echo1))
	b.Run("echo1,args=100", bench(100, echo1))
	b.Run("echo1,args=1000", bench(1000, echo1))
	b.Run("echo1,args=10000", bench(10000, echo1))
	b.Run("echo3,args=10", bench(10, echo3))
	b.Run("echo3,args=100", bench(100, echo3))
	b.Run("echo3,args=1000", bench(1000, echo3))
	b.Run("echo3,args=10000", bench(10000, echo3))
}
