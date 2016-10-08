package insert

import (
	"math/rand"
	"sort"
)

type node struct {
	next *node
	val  int
}

func insertList(n int) *node {
	rand.Seed(0)

	var root *node
	for i := 0; i < n; i++ {
		x := rand.Intn(n)

		var prev *node
		var curr = root
		for curr != nil && curr.val < x {
			prev, curr = curr, curr.next
		}

		nn := &node{next: curr, val: x}

		if prev != nil {
			prev.next = nn
		} else {
			root = nn
		}
	}

	return root
}

func deleteList(root *node, size int) {

	for root != nil {
		i := rand.Intn(size)
		var prev *node
		var curr = root
		for j := 0; j < i; j++ {
			prev, curr = curr, curr.next
		}
		if prev != nil {
			prev.next = curr.next
		} else {
			root = curr.next
		}

		size--
	}
}

func insertSlice(n int) []int {
	rand.Seed(0)

	var s []int
	for i := 0; i < n; i++ {
		x := rand.Intn(n)
		idx := sort.Search(len(s), func(j int) bool { return s[j] >= x })
		s = append(s, 0)
		copy(s[idx+1:], s[idx:])
		s[idx] = x
	}

	return s
}

func deleteSlice(s []int) {
	for len(s) > 0 {
		i := rand.Intn(len(s))
		s = s[:i+copy(s[i:], s[i+1:])]
	}
}
