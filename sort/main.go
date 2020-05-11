package main

import (
	"fmt"
)

func main() {
	a := []int{12, 6, 4534, 523, 12, 1, 2, 312, 23, 56, 7, 8, 9}

	mergesort(a, 0, len(a))
	fmt.Println(a)

	heapsort(a, len(a))
	fmt.Println(a)
}
