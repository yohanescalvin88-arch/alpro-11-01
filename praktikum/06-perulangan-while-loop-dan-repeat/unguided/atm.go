package main

import "fmt"

func main() {
    const passwordBenar = "1234"
    saldo := 10000000
    kesempatan := 3
    login := false

    // Proses login
    for kesempatan > 0 {
        var password string

        fmt.Print("Masukkan password ATM: ")
        fmt.Scan(&password)

        if password == passwordBenar {
            fmt.Println("Login berhasil!")
            login = true
            break
        } else {
            kesempatan--

            if kesempatan > 0 {
                fmt.Println("Password salah. Silakan coba lagi.")
            } else {
                fmt.Println("ATM terblokir. Program selesai.")
            }
        }
    }

    // Menu ATM
    for login {
        fmt.Println("\nMenu ATM:")
        fmt.Println("1. Lihat Saldo")
        fmt.Println("2. Tarik Uang")
        fmt.Println("3. Keluar")
        fmt.Print("Pilih: ")

        var pilihan int
        fmt.Scan(&pilihan)

        switch pilihan {
        case 1:
            fmt.Printf("Saldo Anda: Rp %d\n", saldo)

        case 2:
            var jumlah int

            fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp ")
            fmt.Scan(&jumlah)

            if jumlah <= 0 {
                fmt.Println("Jumlah penarikan harus lebih dari 0.")
            } else if jumlah > saldo {
                fmt.Println("Saldo tidak mencukupi.")
            } else {
                saldo -= jumlah
                fmt.Printf("Anda berhasil menarik uang sebesar Rp %d\n", jumlah)
                fmt.Printf("Saldo tersisa: Rp %d\n", saldo)
            }

        case 3:
            fmt.Println("Keluar dari sistem. Terima kasih!")
            login = false

        default:
            fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
        }
    }
}
