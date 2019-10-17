package rbtree

import (
	"fmt"
	"runtime"
	"strconv"
)

var (
	Debug bool
)

const (
	BLACK = 1
	RED   = 2
)

type Node struct {
	Color  int
	Value  int
	LeftC  *Node
	RightC *Node
	Parent *Node
}

type RbTree struct {
	Root *Node
}

func NewRbTree() *RbTree {
	root := &Node{}
	return &RbTree{root}
}

func (t *RbTree) Debug() {
	if Debug {
		_, _, line, _ := runtime.Caller(1)
		fmt.Println("Debug:", line)
		fmt.Println(t)
	}
}

func (t *RbTree) Put(a int) {
	if t.Root.Color == 0 {
		t.Root.Color = BLACK
		t.Root.Value = a
		return
	}
	var index *Node = t.Root
	var father *Node
	var left bool
	for index != nil {
		father = index
		if a <= index.Value {
			index = index.LeftC
			left = true
		} else {
			index = index.RightC
			left = false
		}
	}
	index = &Node{Color: RED, Value: a, Parent: father}
	if left {
		father.LeftC = index
	} else {
		father.RightC = index
	}
	t.Debug()
	t.fix(index)
}

func (t *RbTree) String() string {
	ret := "Value:Color[:Parent]\n"
	var index []*Node
	var nexts []*Node = []*Node{t.Root}
	for len(nexts) != 0 {
		index = []*Node{}
		for _, v := range nexts {
			if v == nil {
				ret += " nil:1"
			} else {
				ret += fmt.Sprintf(" %d:%d", v.Value, v.Color)
				if v.Parent != nil {
					ret += ":" + strconv.Itoa(v.Parent.Value)
				}
				if v.LeftC != nil {
					index = append(index, v.LeftC)
				} else {
					index = append(index, nil)
				}
				if v.RightC != nil {
					index = append(index, v.RightC)
				} else {
					index = append(index, nil)
				}
			}
		}
		ret += " \n"
		nexts = append([]*Node{}, index...)
	}
	return ret
}

func (t *RbTree) fix(x *Node) {
	var parent, uncle, grandp *Node
	// if x is root; return
	if x.Parent == nil {
		x.Color = BLACK
		t.Debug()
		return
	}

	// if parent.Color == black; return
	if x.Parent.Color == BLACK {
		t.Debug()
		return
	}

	// assign var
	parent = x.Parent
	grandp = parent.Parent
	if grandp.LeftC == parent {
		uncle = grandp.RightC
	} else {
		uncle = grandp.LeftC
	}

	// now parent is RED, gandp is black.
	// when uncle is red
	if uncle != nil && uncle.Color == RED {
		grandp.Color = RED
		uncle.Color = BLACK
		parent.Color = BLACK
		t.fix(grandp)
		t.Debug()
		return
	}
	// now parent is RED, uncle and grandp is BLACK,
	// adjust x and parent.
	if (parent.LeftC == x) && grandp.RightC == parent {
		t.rightRotation(parent)
		parent.Color = BLACK
		t.Debug()
	} else if (parent.RightC == x) && (grandp.LeftC == parent) {
		t.leftRotation(parent)
		parent.Color = BLACK
		t.Debug()
	} else {
		x.Color = BLACK
		t.Debug()
	}

	if grandp.LeftC == parent {
		t.rightRotation(grandp)
		t.Debug()
	} else {
		t.leftRotation(grandp)
		t.Debug()
	}
	t.fix(parent)
}

func (t *RbTree) leftRotation(x *Node) {
	R := x.RightC
	if R == nil {
		panic("leftRotation error")
	}
	parent := x.Parent
	RL := R.LeftC
	R.LeftC = x
	x.Parent = R
	x.RightC = RL
	if RL != nil {
		RL.Parent = x
	}
	if parent == nil {
		t.Root = R
	} else if parent.LeftC == x {
		parent.LeftC = R
	} else {
		parent.RightC = R
	}
	R.Parent = parent
}

func (t *RbTree) rightRotation(x *Node) {
	L := x.LeftC
	if L == nil {
		panic("rightRotation error")
	}
	parent := x.Parent
	LR := L.RightC
	L.RightC = x
	x.Parent = L
	x.LeftC = LR
	if LR != nil {
		LR.Parent = x
	}
	if parent == nil {
		t.Root = L
	} else if parent.LeftC == x {
		parent.LeftC = L
	} else {
		parent.RightC = L
	}
	L.Parent = parent
}
