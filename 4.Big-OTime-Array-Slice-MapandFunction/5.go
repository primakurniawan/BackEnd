package main

import (
	"fmt"
	"sort"
)

func PairSum(arr []int, target int) []int {

	// your code here
	sortedArray := sort.IntSlice(arr)
	low := 0
	high := len(arr) - 1

	for low < high {

		if sortedArray[low]+sortedArray[high] == target {
			return []int{low, high}
		}
		if sortedArray[low]+sortedArray[high] > target {
			high--
		} else {
			low++
		}
	}
	return []int{0}
}

func main() {

	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]

	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11)) // [0, 2]

	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12)) // [2, 3]

	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10)) // [1, 2]

	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6)) // [0, 1]

}
