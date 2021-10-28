package main

import (
	"fmt"
	"strconv"
	"strings"
)

func munculSekali(angka string) []int {

	// your code here
	var list map[int]int
	list = map[int]int{}
	var sliceOnce []int
	intSlice := strings.Split(angka, "")
	for _, v := range intSlice {
		i, _ := strconv.Atoi(v)
		list[i]++
	}
	for i, _ := range list {
		if list[i] == 1 {
			sliceOnce = append(sliceOnce, i)
		}
	}
	return sliceOnce
}

func main() {

	fmt.Println(munculSekali("1234123")) // [4]

	fmt.Println(munculSekali("76523752")) // [6, 3]

	fmt.Println(munculSekali("12345")) // [1 2 3 4 5]

	fmt.Println(munculSekali("1122334455")) // []

	fmt.Println(munculSekali("0872504")) // [8 7 2 5 4]

	try := 2 ^ 7 ^ 3
	fmt.Println(try)
}
