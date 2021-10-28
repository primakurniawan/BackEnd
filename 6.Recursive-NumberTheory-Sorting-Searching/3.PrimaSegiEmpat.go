package main

import "fmt"

func primaSegiEmpat(high, wide, start int) {

	// your code here
	var index int
	var valArr []int

	for index <= high*wide {
		start++

		if isPrime(start, 2) == true {
			valArr = append(valArr, start)
			index++
		}
		if index == high*wide {
			break
		}
	}

	index = 0
	for i := 0; i < high; i++ {
		for j := 0; j < wide; j++ {
			fmt.Print(valArr[index], " ")
			index++
		}
		fmt.Println("\n")
	}

	sumArr := sumArr(valArr)

	fmt.Println(sumArr)

}

func sumArr(arr []int) int {
	var sum int = 0

	for _, v := range arr {
		sum += v
	}
	return sum
}

func main() {

	primaSegiEmpat(2, 3, 13)

	/*

	   17 19

	   23 29

	   31 37

	   156

	*/

	primaSegiEmpat(5, 2, 1)

	/*

	   2  3  5  7 11

	   13 17 19 23 29

	   129

	*/

}
