package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Bilangan harus positif")
		return
	}

	jumlah := 0

	for i := 1; i <= n; i += 2 {
		fmt.Println(i, " ")
		jumlah += i
	}
	
	fmt.Println(jumlah)
}
