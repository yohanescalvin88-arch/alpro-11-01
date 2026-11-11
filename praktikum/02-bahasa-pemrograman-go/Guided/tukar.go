package main

import "fmt"

func main() {
	var a, b int

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
}

	