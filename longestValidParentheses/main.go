package main

import (
	"fmt"
	"strconv"
)

var debug = false

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	queenMap := map[string]bool{}
	for _, queen := range queens {
		x, y := queen[0], queen[1]
		queenMap[strconv.Itoa(x)+strconv.Itoa(y)] = true
	}
	dirictions := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0},
		[]int{1, 1}, []int{1, -1}, []int{-1, 1}, []int{-1, -1}}
	ret := [][]int{}
	for _, diriction := range dirictions {
		if diriction[0] == diriction[1] && diriction[0] == 1 {
			debug = true
		}
		x, y := king[0], king[1]
		for i := 0; i < 7; i++ {
			x, y = x+diriction[0], y+diriction[1]
			if debug {
				fmt.Println(x, y)
			}
			if x < 0 || x > 7 || y < 0 || y > 7 {
				break
			}
			if queenMap[strconv.Itoa(x)+strconv.Itoa(y)] {
				ret = append(ret, []int{x, y})
				break
			}
		}
	}
	return ret
}
func main() {
	fmt.Println(queensAttacktheKing([][]int{[]int{0, 1}, []int{1, 0}, []int{4, 0}, []int{0, 4}, []int{3, 3}, []int{2, 4}}, []int{0, 0}))
}
