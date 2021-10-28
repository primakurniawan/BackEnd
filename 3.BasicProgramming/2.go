package main

import "fmt"

func main() {
	var nilai int
	var deskripsi string

	fmt.Scanf("%d", &nilai)

	if nilai >=80 && nilai <=100 {
		deskripsi = "A"
	} else if nilai >=65 && nilai <=79 {
		deskripsi = "B"
	}else if nilai >=50 && nilai <=64 {
		deskripsi = "C"
	} else if nilai >=35 && nilai <=49 {
		deskripsi = "D"
	} else if nilai >=0 && nilai <=34 {
		deskripsi = "A"
	} else {
		deskripsi = "Nilai Invalid"
	}

	fmt.Printf("%s", deskripsi)
}