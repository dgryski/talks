package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type s struct {
	b bool
	i int
}

func main() {

	slice := make([]s, 20)

	// START OMIT

	slice14 := &slice[14]

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	off14 := sh.Data + 14*unsafe.Sizeof(slice[0])

	// END OMIT

	fmt.Printf("slice14  = %x\n", uintptr(unsafe.Pointer(slice14)))
	fmt.Printf("off14    = %x\n", off14)
}
