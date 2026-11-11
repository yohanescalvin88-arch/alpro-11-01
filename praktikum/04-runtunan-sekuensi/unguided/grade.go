package main

import (
	"fmt"
)

func main() {
	var nama string
	var nilai float64
	var keterangan string

	fmt.Scan(&nama)
	fmt.Scan(&nilai)

	switch {
	case nilai >= 90 && nilai <= 100:
		keterangan = "Lee Haechan mendapatkan nilai A"
	case nilai >= 80 && nilai < 90:
		keterangan = "Lee Haechan mendapatkan nilai B"
	case nilai >= 70 && nilai < 80:
		keterangan = "Lee Haechan mendapatkan nilai C"
	case nilai >= 60 && nilai < 70:
		keterangan = "Lee Haechan mendapatkan nilai D"
	case nilai >= 0 && nilai < 60:
		keterangan = "Lee Haechan mendapatkan nilai F"
	}
	fmt.Println(keterangan)
}