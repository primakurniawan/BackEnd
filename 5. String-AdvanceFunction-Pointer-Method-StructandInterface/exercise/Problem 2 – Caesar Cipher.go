package main

import "fmt"

func caesar(offset int, input string) string {

	// your code here
	var output string
	for _, v := range input {
		n := (rune(v) + rune(offset) - 97) % (26)
		output += string(rune(97 + n))
	}
	return output

}

func main() {

	fmt.Println(caesar(3, "abc")) // def

	fmt.Println(caesar(2, "alta")) // cnvc

	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi

	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) //bcdefghijklmnopqrstuvwxyza

	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) //mnopqrstuvwxyzabcdefghijkl

}
