package main

import "fmt"

func main() {
    var n int

    fmt.Print("Masukkan tahun ke-n: ")
    fmt.Scan(&n)

    biaya := float64(10000000)
    tahun := 1

    for tahun < n {
        biaya = biaya * 1.06
        tahun++
    }
    fmt.Printf("%.0f\n", biaya)
}
