package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) {

	// your code here
	sort.SliceStable(dragonHead, func(i, j int) bool {
		return dragonHead[i] < dragonHead[j]
	})
	sort.SliceStable(knightHeight, func(i, j int) bool {
		return knightHeight[i] < knightHeight[j]
	})

	curDragonHead := 0
	curKnightHeight := 0
	minimumKnightHeightTotal := 0
	for curDragonHead < len(dragonHead) && curKnightHeight < len(knightHeight) {
		if dragonHead[curDragonHead] <= knightHeight[curKnightHeight] {
			minimumKnightHeightTotal += knightHeight[curKnightHeight]
			curDragonHead++
		}
		curKnightHeight++
	}
	if curDragonHead == len(dragonHead) {
		fmt.Println(minimumKnightHeightTotal)
	} else {
		fmt.Println("knight falls")
	}

}

func main() {

	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4}) // 11

	DragonOfLoowater([]int{5, 10}, []int{5}) // knight fall

	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall

	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10

}
