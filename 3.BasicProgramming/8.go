package main

import "fmt"

func cetakTablePerkalian(number int) {

	// Process: Your Solution Code Here
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Printf("%d ",i*j)
		}
		fmt.Printf("\n")
	}

}

func main() {
	cetakTablePerkalian(9)
}