package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	sepuluhribu := x / 10000
	sisa := x % 10000

	limaribu := sisa / 5000
	sisa = sisa % 5000

	seribu := sisa / 1000

	fmt.Println(sepuluhribu, limaribu, seribu)
}