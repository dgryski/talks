package main

import (
	"strconv"
	"testing"
)

var sink int

func benchmarkGrow(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		var s []int
		for j := 0; j < n; j++ {
			s = append(s, j)
		}
		sink += len(s)
	}
}

func benchmarkAllocate(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		s := make([]int, n)
		for j := 0; j < n; j++ {
			s = append(s, j)
		}
		sink += len(s)
	}
}

func BenchmarkGrow(b *testing.B) {
	for i := 10; i < 200; i += 10 {
		b.Run(strconv.Itoa(i), func(b *testing.B) { benchmarkGrow(b, i) })
	}
}

func BenchmarkAllocate(b *testing.B) {
	for i := 10; i < 2000; i += 10 {
		b.Run(strconv.Itoa(i), func(b *testing.B) { benchmarkAllocate(b, i) })
	}
}
