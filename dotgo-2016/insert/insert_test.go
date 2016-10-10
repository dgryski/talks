package insert

import (
	"strconv"
	"testing"
)

func TestSlice(t *testing.T) {
	v := insertSlice(10)
	t.Log(v)
	deleteSlice(v)
}

func benchmarkSlice(b *testing.B, n int) {
	for bn := 0; bn < b.N; bn++ {
		s := insertSlice(n)
		deleteSlice(s)
	}
}

func TestList(t *testing.T) {
	const size = 10
	v := insertList(size)
	for n := v; n != nil; n = n.next {
		t.Log(n.val)
	}
	deleteList(v, size)
}

func benchmarkList(b *testing.B, n int) {
	for bn := 0; bn < b.N; bn++ {
		r := insertList(n)
		deleteList(r, n)
	}
}

func BenchmarkSlice(b *testing.B) {
	sizes := []int{10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000, 20000, 50000, 100000, 200000, 500000}
	for _, n := range sizes {
		b.Run(strconv.Itoa(n), func(b *testing.B) { benchmarkSlice(b, n) })
	}
}

func BenchmarkList(b *testing.B) {
	sizes := []int{10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000, 20000, 50000, 100000, 200000, 500000}
	for _, n := range sizes {
		b.Run(strconv.Itoa(n), func(b *testing.B) { benchmarkList(b, n) })
	}
}
