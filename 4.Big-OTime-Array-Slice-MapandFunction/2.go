// Problem 2 – Fast Exponentiation

package main

import "fmt"

func pow(x, n int) int {

	// your code here
	if n < 0 {
		return pow(1/x, -n)
	} else if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else if n%2 == 0 {
		return pow(x*x, n/2)
	} else if n%2 != 0 {
		return x * pow(x*x, (n-1)/2)
	}

	return 0

}

func main() {

	fmt.Println(pow(2, 3)) // 8

	fmt.Println(pow(7, 2)) // 49

	fmt.Println(pow(10, 5)) // 100000

	fmt.Println(pow(17, 6)) // 24137569

	fmt.Println(pow(5, 3)) // 125

}
