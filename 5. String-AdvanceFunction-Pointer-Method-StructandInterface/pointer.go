package main

import "fmt"

func main() {
	var name string = "prime"
	var nameAddres *string = &name

	fmt.Println("name (value): ", name)
	fmt.Println("nameAddress (address): ", &name)
	fmt.Println("nameAddress (value): ", *nameAddres)
	fmt.Println("nameAddress (address): ", nameAddres)
}
