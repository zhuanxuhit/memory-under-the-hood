package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arrays := [2]int{1, 2}
	slices := arrays[:2]
	fmt.Println(arrays, slices)
	slices = append(slices, 3)
	fmt.Println(arrays, slices)
	slices[0] = 0
	fmt.Println(arrays, slices)

	intChan := make(chan int)
	go func() {
		intChan <- 1
	}()
	<-intChan
	fmt.Println(unsafe.Sizeof(intChan))

	type Better struct {
		Exists        bool
		PreciseNumber float32
		JustAnInt     int64
	}

	type AlittleInefficient struct {
		PreciseNumber float32
		JustAnInt     int64
		Exists        bool
	}
	a := AlittleInefficient{}
	b := Better{}
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
}
