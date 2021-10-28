package main

import "fmt"

func primeNumber(number int) bool {

	// your code here
	var current int
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			current++
		}
		if current > 2 {
			return false
		}
	}
	return true

}

func main() {
	fmt.Println(primeNumber(19))
}