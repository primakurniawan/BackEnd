package main

import "fmt"

func isPrime(n, i int) bool {
	// Base cases
	if n <= 2 {
		if n == 2 {
			return true
		} else {
			return false
		}
	}
	if n%i == 0 {
		return false
	}
	if i*i > n {
		return true
	}

	// Check for next divisor
	return isPrime(n, i+1)
}

func primeX(number int) int {

	// your code here
	n := 0
	i := 1
	for n <= number {
		if isPrime(i, 2) == true {
			n++
		}
		if n == number {
			return i
		}
		i++
	}
	return -1
}

func main() {
	fmt.Println(primeX(1)) // 2

	fmt.Println(primeX(5)) // 11

	fmt.Println(primeX(8)) // 19

	fmt.Println(primeX(9)) // 23

	fmt.Println(primeX(10)) // 29
}
