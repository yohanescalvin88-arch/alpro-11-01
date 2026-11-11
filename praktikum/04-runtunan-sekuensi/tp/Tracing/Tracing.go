package main

import "fmt"

func main() {
	x := 10
	y := 5
	z := 15
	result := 0

	if x > 5 {
		if y < 10 {
			result = x + y
		} else {
			result = x - y
		}
	}

	if z > 10 && x == 10 {
		result += z
	} else {
		result = z - x
	}

	if x == 10 || y > 10 {
		result += 5
	} else if y == 5 && z > 10 { 
		result -= 5 
	} else {
		result *= 2 
	}

	 if !(x < 15 && y < 10) { 
		 result += 10 
	 } else {
		 result -= 10
	 }

	 fmt.Println("Nilai akhir result:", result) 
}


1. Berapa nilai akhir dari variabel result setelah semua pernyataan kondisi dieksekusi?
25
2. Apa output yang dihasilkan oleh program?
Nilai akhir result: 25
3. Tuliskan langkah-langkah alur eksekusi program berdasarkan kondisi yang diberikan:
Kondisi 1: Apakah kondisi x > 5 benar? Jika ya, apa yang terjadi selanjutnya?
- Karena x = 10, maka x > 5 adalah benar.
Kondisi 2: Apakah kondisi z > 10 && x == 10 benar? Bagaimana hal ini memengaruhi nilai result? 
- karena z = 15 dan x = 10, maka pernyataan tersebut benar (True). program mengesekusi blok result += z
Kondisi 3: Apakah salah satu dari kondisi x == 10 || y > 10 benar? Apa yang terjadi?
- karena x = 10 (True) atau y > 10 (False). jadi secara keseluruhan (True). program mengesekusi perintah result += 5