package main

import "fmt"

type student struct {
	name string

	nameEncode string

	score int
}

type Chiper interface {
	Encode() string

	Decode() string
}

func (s *student) Encode() string {

	var nameEncode = ""

	// your code here
	for _, v := range s.name {
		n := (25 - (rune(v) - 97)) % (26)
		nameEncode += string(rune(97 + n))
	}

	return nameEncode

}

func (s *student) Decode() string {

	var nameDecode = ""

	// your code here
	for _, v := range s.nameEncode {
		n := (25 - (rune(v) - 97)) % (26)
		nameDecode += string(rune(97 + n))
	}

	return nameDecode

}

func main() {

	var menu int

	var a = student{}

	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")

	fmt.Scan(&menu)

	if menu == 1 {

		fmt.Print("\nInput Student’s Name : ")

		fmt.Scan(&a.name)

		fmt.Print("\nEncode of Student’s Name " + a.name + " is : " + c.Encode())

	} else if menu == 2 {

		fmt.Print("\nInput Student’s Encode Name : ")

		fmt.Scan(&a.nameEncode)

		fmt.Print("\nDecode of Student’s Name " + a.nameEncode + " is : " + c.Decode())

	} else {

		fmt.Println("Wrong input name menu")

	}

}
