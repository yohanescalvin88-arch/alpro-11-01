package main

import "fmt"

func main() {
	var m int
	fmt.Scan(&m)

	var jumlahLulus int = 0

	for i := 0; i < m; i++ {
		var p, v, ga, fs, ci, tc int
		fmt.Scan(&p, &v, &ga, &fs, &ci, &tc)

		total := float64(p + v + ga + fs + ci + tc)
		rataRata := total / 6.0

		if rataRata >= 3.33 {
			fmt.Println("Mendapat Sertifikat")
			jumlahLulus++
		} else {
			fmt.Println("Tidak Mendapat Sertifikat")
		}
	}
	fmt.Println(jumlahLulus)
}