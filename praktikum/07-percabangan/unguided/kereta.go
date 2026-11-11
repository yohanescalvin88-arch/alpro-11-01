package main

import (
	"fmt"
)

func main() {
	var kelas string
	var usia int 
	var hari string

	fmt.Scan(&kelas, &usia, &hari)

	var hargadasar float64

	switch kelas {
	case "Ekonomi":
		hargadasar = 100000
	case "Bisnis":
		hargadasar = 200000
	case "Eksekutif":
		hargadasar = 350000
	default:
		fmt.Println("kelas tidak valid")
		return
	}

	hargatiket := hargadasar

	if hari == "Sabtu" || hari == "Minggu" {
		hargatiket = hargatiket + (0.20 * hargatiket)
	}
	
	if usia < 3 {
		hargatiket = 0
	} else if usia >= 3 && usia <= 12 {
		hargatiket = hargatiket - (0.50 * hargatiket)
	} else if usia >= 60 {
		hargatiket = hargatiket - (0.30 * hargatiket)
	}

	fmt.Println(hargatiket)






	

}