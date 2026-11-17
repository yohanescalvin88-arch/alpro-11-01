package main

import "fmt"

func main() {
	var totalHari, tahun, bulan, minggu, sisaHari int

	fmt.Scan(&totalHari)

	tahun = totalHari / 360
	sisaHari = totalHari % 360

	bulan = sisaHari / 30
	sisaHari = sisaHari % 30

	minggu = sisaHari / 7
	sisaHari = sisaHari % 7

	fmt.Println(tahun)
	fmt.Println(bulan)
	fmt.Println(minggu)
	fmt.Println(sisaHari)
}