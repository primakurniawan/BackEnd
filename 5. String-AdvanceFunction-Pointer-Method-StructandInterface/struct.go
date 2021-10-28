package main

import "fmt"

type Student struct {
	firstname   string
	lastname    string
	gpa         float32
	isGraduated bool
}

func (e Student) fullname() string {
	return e.firstname + " " + e.lastname
}

func main() {
	var prime Student = Student{
		firstname:   "prima",
		lastname:    "kurniawan",
		gpa:         3.45,
		isGraduated: false,
	}
	fmt.Println(prime.firstname)
	fmt.Println(prime.gpa)
	fmt.Println(prime.isGraduated)
	fmt.Println(prime.fullname())
}
