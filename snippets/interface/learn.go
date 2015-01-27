package main

import (
	"fmt"
)

type S struct {
	i int
}

func (p *S) Get() int {
	return p.i
}
func (p *S) Put(v int) {
	p.i = v
}

type I interface {
	Get() int
	Put(int)
}

func f(p I) {
	switch t := p.(type) {
	case *S:
	case *R:
	case S:
	case R:
	default:
	}
}

type R struct {
	i int
}

func (p *R) Get() int {
	return p.i
}

func (p *R) Put(v int) {
	p.i = v
}

func main() {
	var s S
	f(&s)
	f(&s)
	f(&s)
}
