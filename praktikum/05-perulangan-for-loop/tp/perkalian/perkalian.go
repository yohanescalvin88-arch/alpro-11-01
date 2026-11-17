package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Masukkan bilangan a: ")
	fmt.Scan(&a)

	fmt.Print("masukkan batas perkalian b:")
	fmt.Scan(&b)

	if a <= 0 || b <= 0 {
		fmt.Println("Bilangan harus positif")
		return
	}

	for i := 1; i <= b; i++ {
		fmt.Printf("%d x %d = %d\n", a, i, a*i)
	}
}