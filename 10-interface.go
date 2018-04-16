package main

import "fmt"

type I interface {
	Get() int
	Set(int)
}

type R struct {
	i int
}
type S struct {
	i int
}

func (p *S) Get() int {
	return p.i
}

func (p *S) Set(v int) {
	p.i=v
}
func (p *R) Get() int {
	return p.i
}

func (p *R) Set(v int) {
	p.i = v
}

func f(p I) {
	p.Set(3)
	fmt.Println(p.Get())
	switch  p.(type) {
	case *S:
		fmt.Println("sss")
	case *R:
		fmt.Println("rrr")
	default:
		fmt.Println("default")
	}
	// 判断 p 是否实现了 I 接口
	if t, ok := p.(I) ; ok {
		fmt.Println(t, ok)
	}
}

func main() {
	var s S; f(&s)
}