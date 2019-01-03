package main

import (
	"fmt"
	"unsafe"
)

func testM(m map[int64]int64) {
	m[1] = 3
}

func main() {
	m := make(map[int64]int64)
	m[1] = 2
	testM(m)
	// m 本身就是一个指针
	fmt.Print(unsafe.Sizeof(m),"\n")
	fmt.Printf("%v\n",m)
}
