package main

import (
	"fmt"
)

func main() {
	var penghasilan float64
	var pajak float64

	fmt.Print("Masukkan penghasilan (dalam juta): ")
	fmt.Scan(&penghasilan)
	
	if penghasilan <= 50 {
		pajak = 0.05 * penghasilan
		
	}
}