package main

import "fmt"

func pangkat(base, pangkat int) int {

	// your code here
	var hasil int = 1
	for i := 1; i <= pangkat; i++ {
		hasil*=base
	}
	return hasil
}

func main() {
	fmt.Println(pangkat(3, 3))
}