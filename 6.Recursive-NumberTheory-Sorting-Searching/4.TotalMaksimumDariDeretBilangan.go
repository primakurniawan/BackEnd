package main

import "fmt"

func sumSubArr(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func MaxSequence(arr []int) int {

	// your code here
	var maxSumSubArr int = arr[0]
	for i := range arr {
		for j := len(arr) - 1; j > i; j-- {
			sumSubArr := sumSubArr(arr[i:j])
			if sumSubArr > maxSumSubArr {
				maxSumSubArr = sumSubArr
			}
		}
	}
	return maxSumSubArr
}

func main() {

	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6})) // 7

	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3})) // 7

	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6})) // 8

	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6})) // 12

}
