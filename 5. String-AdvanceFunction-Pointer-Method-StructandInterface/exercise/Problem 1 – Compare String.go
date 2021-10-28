// Problem 1 – Compare String

package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {

	// your code here
	var contain bool
	if len(a) >= len(b) {
		contain = strings.Contains(a, b)
	} else if len(b) >= len(a) {
		contain = strings.Contains(b, a)
	}

	if contain == true && len(a) >= len(b) {
		return b
	} else if contain == true && len(b) >= len(a) {
		return a
	}

	return ""

}

func main() {

	fmt.Println(Compare("AKA", "AKASHI")) // AKA

	fmt.Println(Compare("KANGOORO", "KANG")) // KANG

	fmt.Println(Compare("KI", "KIJANG")) // KI

	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU

	fmt.Println(Compare("ILALANG", "ILA")) // ILA

}
