package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)

	x, y, z = z, x, y
	
	fmt.Println(x, y, z)
}