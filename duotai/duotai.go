package duotai

import (
	"fmt"
)

type A interface {
	Method() int
	Method1() int
}

type DefaultA struct{}

func (b *DefaultA) Method() int {
	fmt.Println("this is DefaultA.Method.")
	return 0
}
func (b *DefaultA) Method1() int {
	fmt.Println("this is DefaultA.Method1.")
	return 0
}

type B struct {
	A
}

//func (b *B) Method() int {
//	fmt.Println("this is B.Method.")
//	return 1
//}
