
package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Bilangan harus positif")
		return
	}

	fmt.Print("Faktor dari ", n, ": ")

	jumlahFaktor := 0

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
			jumlahFaktor++
		}
	}

	fmt.Println()

	if jumlahFaktor == 2 {
		fmt.Println("Ya")
		fmt.Println(n, "adalah bilangan prima")
	} else {
		fmt.Println("Tidak")
		fmt.Println(n, "bukan bilangan prima")
	}
}
