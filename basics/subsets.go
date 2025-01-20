package main

import "fmt"

func getKSubsets(nums []int, k int) [][]int {

	result := [][]int{}

	var backtrack func(int, []int)
	backtrack = func(pos int, path []int) {
		if len(path) == k {
			arr := make([]int, len(path))
			copy(arr, path)
			result = append(result, arr)
			return
		}

		needed := k - len(path) - 1
		for i := pos; i < len(nums)-needed; i++ {
			// pick
			path = append(path, nums[i])

			// recur
			backtrack(i+1, path)

			// unpick
			path = path[0 : len(path)-1]
		}

	}
	backtrack(0, []int{})
	return result
}

func getKSubsets_2(nums []int, k int) [][]int {

	result := [][]int{}

	var backtrack func(int, []int)
	backtrack = func(pos int, path []int) {
		if len(path) == k {
			arr := make([]int, len(path))
			copy(arr, path)
			result = append(result, arr)
			return
		}

		if pos == len(nums) {
			return
		}

		// pick and recur
		path = append(path, nums[pos])
		backtrack(pos+1, path)

		// unpick and recur
		path = path[0 : len(path)-1]
		backtrack(pos+1, path)
	}
	backtrack(0, []int{})
	return result
}

/*

generate all subsets
1 2 3 4

[[1]]
[1] [2 3 4] (2 23 24 3 34 4 234)

*/

func getAllSubsets(nums []int) [][]int {

	var subsets func(int) [][]int
	subsets = func(start int) [][]int {
		result := [][]int{}
		result = append(result, []int{nums[start]})
		if start == len(nums)-1 {
			return result
		}

		partialSet := subsets(start + 1)
		partialSetPlusThis := [][]int{}
		for _, set := range partialSet {
			result = append(result, set)
			arr := make([]int, len(set))
			copy(arr, set)
			arr = append([]int{nums[start]}, arr...)
			partialSetPlusThis = append(partialSetPlusThis, arr)
		}
		result = append(result, partialSetPlusThis...)
		return result
	}
	return subsets(0)
}

func main() {
	//fmt.Println(getKSubsets_2([]int{1, 2, 3, 4, 5, 6}, 3))
	fmt.Println(getAllSubsets([]int{1, 2, 3}))
}
