package main

import (
	"fmt"
	"unsafe"
	"encoding/json"
	"math"
)

func main() {
	//unsafe.Pointer()
	var definiteArray [8]int64
	fmt.Println(unsafe.Sizeof(definiteArray))
	fmt.Println(unsafe.Sizeof(uintptr(0)))
	//fmt.Println(runtime._MaxMHeapList)

	v , _ := json.Marshal([]int64{11121,12312})
	fmt.Println(string(v))
	fmt.Println(math.Exp2(30)/1024/1024/1024)
	var maps [1024*1024*1024]int32
	maps[0] = 1
}
