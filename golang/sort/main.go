package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"
)

func main() {
	f, err := os.OpenFile("./intarray.log", os.O_RDONLY, 222)
	if err != nil {
		fmt.Println("os.Open: %s", err.Error())
		return
	}
	defer func() {
		err = f.Close()
		if err != nil {
			fmt.Println("f.Close:%v", err)
		}
	}()
	bt := make([]byte, 1000000)
	n, err := f.Read(bt)
	if err != nil {
		fmt.Println("f.Read:%s", err.Error())
		return
	}
	bt = bt[:n]
	var s []int
	json.Unmarshal(bt, &s)
	start := time.Now()
	fast(s)
	fmt.Println(time.Since(start))

	a := append([]int{}, s...)
	start = time.Now()
	sort.Ints(a)
	fmt.Println(time.Since(start))
	for i := 0; i < len(a); i++ {
		if s[i] != a[i] {
			fmt.Println(i, s[i], a[i])
			break
		}
	}

}
func heapSort(a []int) {
	L := len(a)
	if L < 12 {
		maopao(a)
		return
	}
	for i := (L - 1) / 2; i >= 0; i-- {
		shiftDown(a, i, L)
	}
	for i := L - 1; i >= 0; i-- {
		a[i], a[0] = a[0], a[i]
		if i == 1 {
			break
		}
		shiftDown(a, 0, i)
	}
}
func shiftDown(a []int, i, end int) {
	if 2*i+1 >= end {
		return
	}
	var max int = a[i]
	var index int = i
	if a[2*i+1] > max {
		max = a[2*i+1]
		index = 2*i + 1
	}
	if 2*i+2 < end && a[2*i+2] > max {
		max = a[2*i+2]
		index = 2*i + 2
	}
	if max != a[i] {
		a[i], a[index] = a[index], a[i]
		shiftDown(a, index, end)
	}
}
func maopao(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i; j++ {
			if j+1 < len(a)-i && a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
func fast(a []int) {
	if len(a) < 12 {
		maopao(a)
		return
	}
	i, j := 1, len(a)-1
	for i < j {
		for i < j && a[0] >= a[i] {
			i++
		}
		for i < j && a[0] <= a[j] {
			j--
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	if a[0] > a[i] {
		a[0], a[i] = a[i], a[0]
	}
	fast(a[:i])
	fast(a[i:])
}
func insert(a []int) {
	if len(a) < 2 {
		return
	}
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}
	}
}
