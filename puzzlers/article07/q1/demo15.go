package main

import "fmt"

func main() {
	// 示例1。
	s1 := make([]int, 5) // 只指定 len, 不指定 cap，默认为 len
	fmt.Printf("The length of s1: %d\n", len(s1)) // 5
	fmt.Printf("The capacity of s1: %d\n", cap(s1)) // 5
	fmt.Printf("The value of s1: %d\n", s1) // [0,0,0,0,0]
	s2 := make([]int, 5, 8) // 指定 len, cap
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)
	fmt.Println()

	// 示例2。
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6] // 底部都公用一个数组，所以cap是数组最大长度
	fmt.Printf("The length of s4: %d\n", len(s4)) // 3
	fmt.Printf("The capacity of s4: %d\n", cap(s4)) // 5
	fmt.Printf("The value of s4: %d\n", s4) //[4 5 6]
	fmt.Println()

	// 示例3。
	s5 := s4[:cap(s4)]
	fmt.Printf("The length of s5: %d\n", len(s5)) // 5
	fmt.Printf("The capacity of s5: %d\n", cap(s5)) // 5
	fmt.Printf("The value of s5: %d\n", s5) //[4 5 6 7 8]
}