package main

import "fmt"

/*
[
["0","1","1","0","0","1","0","1","0","1"],
["0","0","1","0","1","0","1","0","1","0"],
["1","0","0","0","0","1","0","1","1","0"],
["0","1","1","1","1","1","1","0","1","0"],
["0","0","1","1","1","1","1","1","1","0"],
["1","1","0","1","0","1","1","1","1","0"],
["0","0","0","1","1","0","0","0","1","0"],
["1","1","0","1","1","0","0","1","1","1"],
["0","1","0","1","1","0","1","0","1","1"]]

[
[0 1 1 0 0 1 0 1 0 1]
[0 0 2 0 1 0 1 0 1 0]
[1 0 0 0 0 1 0 1 2 0]
[0 1 1 1 1 2 1 0 3 0]
[0 0 2 2 2 3 2 1 4 0]
[1 1 0 3 0 4 3 2 5 0]
[0 0 0 4 1 0 0 0 6 0]
[1 1 0 5 2 0 0 1 7 1]
[0 2 0 6 3 0 1 0 8 2]]

[0 1 1 0 0 1 0 1 0 1]

[0 0 2 2 2 3 2 1 4 0] = 10
*/
func maxHistogramArea(hist []int) int {

	st := []int{}
	idx := []int{}
	maxArea := -1

	i := 0
	for ; i < len(hist); i++ {
		fmt.Printf("i: %d, st: %v, idx: %v, max: %d\n", i, st, idx, maxArea)
		curr := hist[i]

		if len(st) > 0 && top(st) == curr {
			continue
		}

		right, left := i, i
		for len(st) > 0 && top(st) > curr {
			left = top(idx)
			area := (right - left) * top(st)
			if area > maxArea {
				maxArea = area
			}
			_, st = pop(st)
			_, idx = pop(idx)
		}

		if curr > 0 {
			if len(st) == 0 || curr != top(st) {
				st = push(st, curr)
				idx = push(idx, left)
			}
		} else {
			for len(st) > 0 {
				left = top(idx)
				area := (right - left) * top(st)
				if area > maxArea {
					maxArea = area
				}
				_, st = pop(st)
				_, idx = pop(idx)
			}
		}
	}

	for len(st) > 0 {
		left := top(idx)
		area := (i - left) * top(st)
		if area > maxArea {
			maxArea = area
		}
		_, st = pop(st)
		_, idx = pop(idx)
	}

	return maxArea
}

func pop(st []int) (int, []int) {
	return st[len(st)-1], st[0 : len(st)-1]
}

func push(st []int, val int) []int {
	st = append(st, val)
	return st
}

func top(st []int) int {
	return st[len(st)-1]
}

func main() {
	fmt.Println("Hello world")
	//h := []int{3, 1, 3, 2, 2, 1, 1, 1, 1, 1, 1, 10000, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	h := []int{0, 0, 2, 2, 2, 3, 2, 1, 4, 0}
	fmt.Println(maxHistogramArea(h))
}
