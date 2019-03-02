package main

import (
	"os"
	"testing"
)

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat(os.Args[1:])
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		join(os.Args[1:])
	}
}
