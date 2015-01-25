package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	if s.i+1 > 9 {
		return
	}
	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() int {
	s.i--
	return s.data[s.i]
}

/*
 * 编写一个String方法将栈转化为字符串形式的表达。可以这样的方式打印整个栈:
 * fmt.Printf("My stack %v\n", stack)
 * 栈可以被输出成这样的形式: [0:m] [1:l] [2:k]
 *
 * 根据 Go 文档 fmt.Printf("\%v") 可以打印实现了 Stringer 接口的任何 值(%v)
 * 为了使其工作,需要为类型定义一个 String() 函数
 */
func (s stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}

func main() {
	var s stack
	s.push(50)
	fmt.Printf("stack %v\n", s)
	s.push(24)
	fmt.Printf("stack %v\n", s)
	s.pop()
	fmt.Printf("stack %v\n", s)
	s.pop()
	fmt.Printf("stack %v\n", s)
}
