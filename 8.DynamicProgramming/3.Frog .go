package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}

var result = make(map[int]int)

func diff(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func Frog(jumps []int) int {
	N := len(jumps) - 1
	result[0] = 0
	result[1] = diff(jumps[0], jumps[1])
	for i := 2; i <= N; i++ {
		result[i] = int(math.Min(float64(result[i-2]+diff(jumps[i], jumps[i-2])), float64(result[i-1]+diff(jumps[i], jumps[i-1]))))
	}

	return result[N]
}
