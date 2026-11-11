package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	if bilangan == 0 {
		fmt.Println("nol")
	} else {
		if bilangan > 0 {
			fmt.Println("positif, ")
		} else {
			fmt.Println("negatif, ")
		}

		if bilangan % 2 == 0 {
			fmt.Println("genap")
		} else {
			fmt.Println("ganjil")
		}
	}
}