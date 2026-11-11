package main

import "fmt"

func main() {
	var karakter string
	var countM, countK, countT int
	 
	for {
		fmt.Scan(&karakter)

		if karakter == "M" {
			countM++
		} else if karakter == "K" {
			countK++
		} else if karakter == "T" {
			countT++
		} else {
			return
	}
	fmt.Println("Tipe M: ", countM)
	fmt.Println("Tipe K: ", countK)
	fmt.Println("Tipe T: ", countT)
	}
}