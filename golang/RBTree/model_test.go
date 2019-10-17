package rbtree

import (
	"fmt"
	"testing"
)

func TestPuft(t *testing.T) {
	tem := NewRbTree()
	a := []int{1, 2, 3, 4, 5}
	for _, v := range a {
		if v == 11 {
			Debug = true
		}
		tem.Put(v)
		fmt.Println(tem)
	}
}

/*

func TestString(t *testing.T) {
	tree := &RbTree{}
	lf := &RbTree{Color: 1, Value: 1, Parent: tree}
	rt := &RbTree{Color: 1, Value: 5, Parent: tree}
	tree.Color = 1
	tree.Value = 3

	tree.LeftC = lf
	tree.RightC = rt

	lflf := &RbTree{Color: 2, Value: -1, Parent: lf}
	rtlf := &RbTree{Color: 2, Value: 4, Parent: rt}
	lf.LeftC = lflf
	rt.LeftC = rtlf
	fmt.Println(tree)
}
func TestLeftRotation(t *testing.T) {
	root := &RbTree{Value: 9, Color: BLACK}
	l := &RbTree{Value: 7, Color: RED, Parent: root}
	lr := &RbTree{Value: 8, Color: RED, Parent: l}
	root.LeftC = l
	l.RightC = lr
	l.leftRotation(l)
	fmt.Println(root)
	root.rightRotation(root)
	fmt.Println(lr)
}
*/
