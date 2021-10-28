package main

import (
	"fmt"
)

func main() {
	// Array
	// var intArray [5]int = [5]int{1, 3, 6, 8, 9}
	
	// for i := 0; i < len(intArray); i++ {
	// 	fmt.Printf("index: %d => value: %d\n", i, intArray[i])
	// }
	
	// for i, el := range intArray {
	// 	fmt.Printf("index: %d => value: %d\n", i, el)
	// }

	// Slice
	// var strSlice []string = []string{"halo", "hi", "hola", "oie"}
	
	// for i := 0; i < len(strSlice); i++ {
	// 	fmt.Printf("index: %d => value: %s\n", i, strSlice[i])
	// }

	// strSlice = append(strSlice, "yossa")
	
	// for i, el := range strSlice {
	// 	fmt.Printf("index: %d => value: %s\n", i, el)
	// }

	// Map
	var mapPrice map[string]int = map[string]int{"ice cream":7000, "meat ball":15000, "fried rice":25000}
	mapPrice["snack"] = 2000
	fmt.Println(mapPrice["snack"])
}