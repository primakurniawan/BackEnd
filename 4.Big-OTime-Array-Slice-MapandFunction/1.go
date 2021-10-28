// Problem 1 – Bilangan Prima

package main

import (
	"fmt"
	"math"
)

 

func primeNumber(number int) bool {

   // your code here
   if number<=1 {
	   return false
   }
	for i := 2; i < int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

 

func main() {

   fmt.Println(primeNumber(1000000007)) // true

   fmt.Println(primeNumber(1500450271)) // true

   fmt.Println(primeNumber(1000000000)) // false

   fmt.Println(primeNumber(10000000019)) // true

   fmt.Println(primeNumber(10000000033)) // true

}