package membership

import (
	"sort"
)

type sliceSet32 []uint32

func (s sliceSet32) search(n uint32) bool {
	for _, v := range s {
		if v == n {
			return true
		}
	}
	return false
}

type sortedSet32 []uint32

func (s sortedSet32) Len() int           { return len(s) }
func (s sortedSet32) Less(i, j int) bool { return s[i] < s[j] }
func (s sortedSet32) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s sortedSet32) linear(n uint32) bool {
	for _, v := range s {
		if v > n {
			return false
		}
		if v == n {
			return true
		}
	}
	return false
}

func (s sortedSet32) binary(n uint32) bool {
	i := sort.Search(len(s), func(i int) bool { return s[i] >= n })
	return i < len(s) && s[i] == n
}

func (s sortedSet32) inlined(n uint32) bool {

	low, high := 0, len(s)

	for low < high {
		mid := low + (high-low)/2
		if s[mid] >= n {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low < len(s) && s[low] == n
}

type mapSet32 map[uint32]bool

func (m mapSet32) search(n uint32) bool {
	_, ok := m[n]
	return ok
}
