package main

import "fmt"

/*
 * Complete the 'firstOccurrence' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. STRING x
 */

func firstOccurrence(s string, x string) int32 {
	// Write your code here
	pos := 0
	check := 0
	for i := 0; i < len(s); i++ {
		if s[i] == x[i-pos] || x[i-pos] == '*' {
			check++
			if check == len(x)-1 {
				return int32(pos)
			}
		} else {
			check = 0
			pos++
		}

	}
	return -1
}

func main() {
	fmt.Println(firstOccurrence("aslhgyantih", "gy*n"))
}
