// Problem 1 – Menghitung Luas Permukaan Tabung
package main

import "fmt"

func main() {
	const phi float32= 22/7
	var T, r float32
	fmt.Scanf("%g", &T)
	fmt.Scanf("%g", &r)

	var L float32 = 2*phi*r*(r+T);

	fmt.Printf("Luas tabung dengan T: %g dan r: %g adalah L: %g ", T, r, L)
}