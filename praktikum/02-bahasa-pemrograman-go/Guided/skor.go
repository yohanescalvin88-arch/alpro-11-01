package main

import "fmt"

func main() {
	var nama string
	
	var matematika, bahasaInggris int

	fmt.Scan(&nama, &matematika, &bahasaInggris)

	Total := matematika + bahasaInggris
	rataRata := Total / 2

	fmt.Println(nama)
	fmt.Println(Total)
	fmt.Println(rataRata)

}