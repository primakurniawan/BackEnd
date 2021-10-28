package main

import "fmt"
 


 

func playWithAsterik(n int) {

   // your code here
   for i := 1; i <= n; i++ {
	for space := 1; space <= n - i; space++ {
	  fmt.Printf("  ")
	}

	for k := 0; k <= 2 * i - 1; k++ {
		fmt.Printf("* ")
		k++
	}
	fmt.Printf("\n");
	}
}

func main(){
	playWithAsterik(5)
}