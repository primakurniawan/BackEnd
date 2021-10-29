package main

import "fmt"

func main() {
	fmt.Println(fibo(0))  // 0
	fmt.Println(fibo(1))  // 1
	fmt.Println(fibo(2))  // 1
	fmt.Println(fibo(3))  // 2
	fmt.Println(fibo(5))  // 5
	fmt.Println(fibo(6))  // 8
	fmt.Println(fibo(7))  // 13
	fmt.Println(fibo(9))  // 34
	fmt.Println(fibo(10)) // 55
}

func fibo(n int) int {
	result := 0
	if n < 2 {
		result = n
	}
	a := 0
	b := 1
	for i := 2; i <= n; i++ {
		result = a + b
		a = result
		b = a
	}
	return result
}
