
package main

import "fmt"

func main() {
    var n int

    fmt.Print("Masukkan jumlah siswa: ")
    fmt.Scan(&n)

    if n < 2 {
        fmt.Println("Jumlah siswa minimal 2.")
        return
    }

    maxInt := int(^uint(0) >> 1)
    terendah := maxInt
    keduaTerendah := maxInt

    namaTerendah := ""
    namaKedua := ""

    i := 1

    for i <= n {
        var nama string
        var nilai int

        fmt.Scan(&nama, &nilai)

        if nilai < terendah {
            keduaTerendah = terendah
            namaKedua = namaTerendah

            terendah = nilai
            namaTerendah = nama
        } else if nilai < keduaTerendah {
            keduaTerendah = nilai
            namaKedua = nama
        }

        i++
    }

    fmt.Println(namaTerendah)
    fmt.Println(namaKedua)
}
