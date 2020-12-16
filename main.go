package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("%d\n", unsafe.Sizeof(1))      // 8
	fmt.Printf("%d\n", unsafe.Sizeof("qweA")) // 16 (длина + указатель)
	type st struct {
		a byte   // 1
		b bool   // 1
		c uint64 // 8
	}
	s := st{}
	fmt.Printf("%d\n", unsafe.Sizeof(s)) // 16
}
