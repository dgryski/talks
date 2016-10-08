package membership

import (
	"sort"
	"strconv"
	"testing"
)

func testSorted(t *testing.T, n int) {
	s := sliceSet32(make([]uint32, n))
	sorted := sortedSet32(make([]uint32, n))
	for i := 0; i < n; i++ {
		needle := xorshift32(uint32(i))
		s[i] = uint32(needle)
		sorted[i] = uint32(needle)
	}

	sort.Sort(sorted)

	for i := 0; i < 1000*n; i++ {
		needle := xorshift32(uint32(i))
		found := s.search(needle)
		s1 := sorted.linear(needle)
		s2 := sorted.binary(needle)
		s3 := sorted.inlined(needle)
		if found != s1 || found != s2 || found != s3 {
			t.Fatalf("mismatch for %v: found=%v s1=%v s2=%v s3=%v", needle, found, s1, s2, s3)
		}
	}
}

func TestSorted(t *testing.T) {
	var sizes = []int{1, 2, 3, 4, 5, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 150, 200}
	for _, n := range sizes {
		t.Run(strconv.Itoa(n), func(t *testing.T) { testSorted(t, n) })
	}
}

var found int

func benchmarkSlice32(b *testing.B, n int) {
	s := sliceSet32(make([]uint32, n))
	for i := 0; i < n; i++ {
		s[i] = xorshift32(uint32(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		needle := xorshift32(uint32(i))
		if s.search(needle) {
			found++
		}
	}
}

func benchmarkSortedLinear32(b *testing.B, n int) {
	s := sortedSet32(make([]uint32, n))
	for i := 0; i < n; i++ {
		s[i] = xorshift32(uint32(i))
	}
	sort.Sort(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		needle := xorshift32(uint32(i))
		if s.linear(needle) {
			found++
		}
	}
}

func benchmarkSortedBinary32(b *testing.B, n int) {
	s := sortedSet32(make([]uint32, n))
	for i := 0; i < n; i++ {
		s[i] = xorshift32(uint32(i))
	}
	sort.Sort(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		needle := xorshift32(uint32(i))
		if s.binary(needle) {
			found++
		}
	}
}

func benchmarkSortedInlined32(b *testing.B, n int) {
	s := sortedSet32(make([]uint32, n))
	for i := 0; i < n; i++ {
		s[i] = xorshift32(uint32(i))
	}
	sort.Sort(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		needle := xorshift32(uint32(i))
		if s.inlined(needle) {
			found++
		}
	}
}

func benchmarkMap32(b *testing.B, n int) {
	m := mapSet32(make(map[uint32]bool))
	for i := 0; i < n; i++ {
		m[xorshift32(uint32(i))] = true
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		needle := xorshift32(uint32(i))
		if m.search(needle) {
			found++
		}
	}
}

func xorshift32(y uint32) uint32 {
	y ^= (y << 13)
	y ^= (y >> 17)
	y ^= (y << 5)
	return y
}

func BenchmarkSlice32(b *testing.B) {
	var sizes = []int{1, 2, 3, 4, 5, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for _, n := range sizes {
		b.Run(strconv.Itoa(n), func(b *testing.B) { benchmarkSlice32(b, n) })
	}
}

func BenchmarkSortedLinear32(b *testing.B) {
	var sizes = []int{1, 2, 3, 4, 5, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for _, n := range sizes {
		b.Run(strconv.Itoa(n), func(b *testing.B) { benchmarkSortedLinear32(b, n) })
	}
}

func BenchmarkSortedBinary32(b *testing.B) {
	var sizes = []int{1, 2, 3, 4, 5, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for _, n := range sizes {
		b.Run(strconv.Itoa(n), func(b *testing.B) { benchmarkSortedBinary32(b, n) })
	}
}

func BenchmarkSortedInlined32(b *testing.B) {
	var sizes = []int{1, 2, 3, 4, 5, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 150, 200, 250, 300, 350, 400, 500, 600, 700, 800, 900, 1000, 1200, 1400, 1600, 1800, 2000, 2200, 2400, 2600, 2800, 3000, 3200, 3400, 3600, 3800, 4000, 4500, 5000, 10000}
	for _, n := range sizes {
		b.Run(strconv.Itoa(n), func(b *testing.B) { benchmarkSortedInlined32(b, n) })
	}
}

func BenchmarkMap32(b *testing.B) {
	var sizes = []int{1, 2, 3, 4, 5, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 150, 200, 250, 300, 350, 400, 500, 600, 700, 800, 900, 1000, 1200, 1400, 1600, 1800, 2000, 2200, 2400, 2600, 2800, 3000, 3200, 3400, 3600, 3800, 4000, 4500, 5000, 10000}
	for _, n := range sizes {
		b.Run(strconv.Itoa(n), func(b *testing.B) { benchmarkMap32(b, n) })
	}
}
