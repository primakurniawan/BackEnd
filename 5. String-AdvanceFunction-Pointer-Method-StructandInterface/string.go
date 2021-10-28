package main

import (
	"fmt"
)

func main() {
	// len
	// var str string = "this is the string"
	// var lengthStr int = len(str)
	// fmt.Println(lengthStr)

	// compare
	// var str1 string = "hello"
	// var str2 string = "hello"
	// var str3 string = "hi"
	// fmt.Println(str1 == str2)
	// fmt.Println(str1 == str3)

	// contains
	// var str1 string = "hello"
	// var str2 string = "hell"
	// var str3 string = "hi"
	// fmt.Println(strings.Contains(str1, str2))
	// fmt.Println(strings.Contains(str1, str3))

	// substring
	// var str string = "cat,dog,ant"
	// fmt.Println(str[4:7])

	// replace
	// var str string = "hello <name>"
	// var newStr string = strings.Replace(str, "<name>", "prime", -1)
	// fmt.Println(newStr)

	// insert
	var str string = "hello"
	var str2 string = "good bye"
	var addedStr string = str + " bokuwa " + str2
	fmt.Println(addedStr)

}
