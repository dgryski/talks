package main

import (
	"fmt"
	"time"
)

func mult3() {
	arr := make([]int32, 128*1024*1024)

	for j := 0; j < 20; j++ {
		for i := 0; i < len(arr); i++ {
			arr[i] *= 3
		}
	}

	// START OMIT
	for _, j := range []int{1, 2, 3, 4, 5, 6, 7, 8,
		9, 10, 11, 12, 13, 14, 15, 16,
		32, 64, 128, 256, 512, 1024} {
		for c := 0; c < 20; c++ {
			t0 := time.Now()
			for i := 0; i < len(arr); i += j {
				arr[i] *= 3
			}
			fmt.Println(j, time.Since(t0))
		}
	}
	// END OMIT
}

func multRand() {
	arr := make([]int32, 128*1024*1024)

	for i := 0; i < len(arr); i++ {
		arr[i] *= 3
	}

	for _, j := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 32, 64, 128, 256, 512, 1024} {
		for c := 0; c < 10; c++ {
			t0 := time.Now()
			for i := 0; i < len(arr); i += j {
				arr[i] *= int32(xorshift32(uint32(i)))
			}
			fmt.Println(j, time.Since(t0))
		}
	}
}

func xorshift32(y uint32) uint32 {
	y ^= (y << 13)
	y ^= (y >> 17)
	y ^= (y << 5)
	return y
}

func main() {
	mult3()
}
