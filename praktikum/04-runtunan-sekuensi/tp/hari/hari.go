package main

import "fmt"

func main() {
	var tahun int
	var bulan string

	fmt.Scan(&tahun)
	fmt.Scan(&bulan)

	var jumlahHari int
	var isValid bool = true

	isKabisat := (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0)


	jumlahHari = 0
	isValid = true

	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		jumlahHari = 31
	case "Apr", "Jun", "Sep", "Nov":
		jumlahHari = 30
	case "Feb":
		if isKabisat {
			jumlahHari = 29
		} else {
			jumlahHari = 28
		}
	default:
		isValid = false
	}

	if isValid {
		fmt.Println(jumlahHari)
	} else {
		fmt.Printf("\"%s\" tidak valid (huruf besar/kecil berbeda dari format yang diminta, seharusnya contoh: \"jan\")\n", bulan)
	}
}