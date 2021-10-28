package main

import "fmt"

// variadic function
func sum(numbers ...int) int {
	var total int
	for _, v := range numbers {
		total += v
	}
	return total
}

// closures
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}

}

func main() {
	// variadic function

	fmt.Println(sum(2, 5, 6, 7))

	// anonymous function
	func() {
		fmt.Println("immediately invoked")
	}()

	call := func() {
		fmt.Println("it needs to be called")

	}
	call()

	// defer
	defer func() {
		fmt.Println("this is call at the last")
	}()
	// closures
	initCounter := counter()
	fmt.Println(initCounter())
	fmt.Println(initCounter())
	fmt.Println(initCounter())

}
