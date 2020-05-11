package main

import (
	"strconv"
)

type Int struct {
	neg   bool
	value []uint32
}

func BigInt(a string) *Int {
	ret := &Int{}
	if a == "" {
		return ret
	}
	if a[0] == '-' {
		ret.neg = true
		a = a[1:]
	}
	l := len(a)
	count := 0
	val := uint32(0)
	for i := l - 1; i >= 0; i-- {
		switch a[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			val = val / 10
			val += uint32(a[i]-'0') * 100000000
			count++
			if count != 9 && i != 0 {
				continue
			}
		default:
			return nil
		}
		ret.value = append(ret.value, val)
		val = 0
		count = 0
	}
	if l%9 == 0 {
		return ret
	}
	for i := l % 9; i < 9; i++ {
		ret.value[len(ret.value)-1] /= 10
	}
	return ret
}

func (a *Int) String() string {
	ret := ""
	if a.neg {
		ret += "-"
	}
	for i := len(a.value) - 1; i >= 0; i-- {
		ret += strconv.Itoa(int(a.value[i]))
	}
	return ret
}

func (a *Int) Less(b *Int) bool {
	if a.neg && b.neg {
		a_ := BigInt(a.String())
		b_ := BigInt(b.String())
		a_.neg, b_.neg = false, false
		return b_.Less(a_)
	}
	if a.neg && !b.neg {
		return true
	}
	if !a.neg && b.neg {
		return false
	}
	// both > 0
	if len(a.value) < len(b.value) {
		return true
	}
	if len(a.value) > len(b.value) {
		return false
	}
	return a.value[len(a.value)-1] < b.value[len(b.value)-1]
}
