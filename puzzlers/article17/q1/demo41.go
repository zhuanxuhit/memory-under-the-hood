package main

import "fmt"

func main() {
	// 示例1。
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	// range 的右侧表达式会初始化时求值一次
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] |= i
		}
	}
	fmt.Println(numbers1)
	fmt.Println()

	// 示例2。
	numbers2 := [...]int{1, 2, 3, 4, 5, 6} // 数组的声明方式
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 { // range右侧是一份拷贝
		// e 是原数组值
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)
	fmt.Println()

	// 示例3。
	numbers3 := []int{1, 2, 3, 4, 5, 6}
	maxIndex3 := len(numbers2) - 1
	for i, e := range numbers3 { // 切片是引用类型，拷贝不拷贝底层数据
		if i == maxIndex3 {
			numbers3[0] += e
		} else {
			numbers3[i+1] += e
		}
	}
	fmt.Println(numbers3)
}
