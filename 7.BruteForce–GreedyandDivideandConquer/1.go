package main

import (
	"fmt"
	"math"
)

func main() {
	calculate(6, 6, 14)
	calculate(10, 30, 38)
	calculate(7, 8, 21)
}

func calculate(a, b, c int) (x, y, z int) {
	q := int(math.Ceil(float64(a) / 2))
	for i := 1; i <= q; i++ {
		for j := 1; j <= q; j++ {
			for k := 1; k <= q; k++ {
				if i+j+k == a {
					if i*j*k == b {
						if i*i+j*j+k*k == c {
							fmt.Println(i, j, k)
							return
						}
					}
				}
			}
		}
	}
	fmt.Println("no solution found")
	return
}
