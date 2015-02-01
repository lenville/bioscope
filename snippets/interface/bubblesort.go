package main

import (
	"fmt"
)

// 定义一个有着若干排序相关的方法的接口类型
type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// 定义用于排序的 slice 的新类型
type Xi []int
type Xs []string

// 实现 Sorter 接口方法
func (p Xi) Len() int {
	return len(p)
}
func (p Xi) Less(i int, j int) bool {
	return p[j] < p[i]
}
func (p Xi) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p Xs) Len() int {
	return len(p)
}
func (p Xs) Less(i int, j int) bool {
	return p[j] < p[i]
}
func (p Xs) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

// 编写作用于 Sorter 接口的通用排序函数
func Sort(x Sorter) {
	for i := 0; i < x.Len()-1; i++ {
		for j := 0; j < x.Len(); j++ {
			if !x.Less(i, j) {
				x.Swap(i, j)
			}
		}
	}
}

func main() {
	ints := Xi{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
	strings := Xs{"nut", "ape", "elephant", "zoo", "go"}

	fmt.Printf("Before: %v\n", ints)
	Sort(ints)
	fmt.Printf("After:  %v\n\n", ints)

	fmt.Printf("Before: %v\n", strings)
	Sort(strings)
	fmt.Printf("After:  %v\n", strings)
}
