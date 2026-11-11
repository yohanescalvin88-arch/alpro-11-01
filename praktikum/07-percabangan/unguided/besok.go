
package main

import "fmt"

func main() {
    var tanggal, bulan, tahun int
    var jumlahHari int
    var kabisat bool

    fmt.Print("Masukkan tanggal, bulan, tahun: ")
    fmt.Scan(&tanggal, &bulan, &tahun)

    if bulan < 1 || bulan > 12 {
        fmt.Println("Tidak valid")
        return
    }


    if tahun%400 == 0 {
        kabisat = true
    } else if tahun%100 == 0 {
        kabisat = false
    } else if tahun%4 == 0 {
        kabisat = true
    } else {
        kabisat = false
    }


    if bulan == 2 {
        if kabisat {
            jumlahHari = 29
        } else {
            jumlahHari = 28
        }
    } else if bulan == 4 || bulan == 6 ||
        bulan == 9 || bulan == 11 {
        jumlahHari = 30
    } else {
        jumlahHari = 31
    }


    if tanggal < 1 || tanggal > jumlahHari {
        fmt.Println("Tidak valid")
        return
    }
    fmt.Println("Valid")

	
    tanggal++

    if tanggal > jumlahHari {
        tanggal = 1
        bulan++

        if bulan > 12 {
            bulan = 1
            tahun++
        }
    }
    fmt.Printf("%d/%d/%d\n", tanggal, bulan, tahun)
}
