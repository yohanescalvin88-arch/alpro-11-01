# <h1 align="center">Laporan Praktikum Modul [2] - [Bahasa Pemrograman go]</h1>
<p align="center">[Variabel, Tipe Data, dan Operasi ] - [109092600013]</p>

## Dasar Teori

### A. [Judul Topik Dasar Teori 1, misal: Bahasa Pemrograman Go]
[Tuliskan penjelasan teori terkait topik ini. Sertakan kutipan/rujukan jika perlu, contoh: Menurut [Yohanes Calvin] ([2026])]
Variabel adalah tempat untuk menyimpan suatu nilai dalam program. Dalam bahasa pemrograman Go, variabel dapat dideklarasikan menggunakan kata kunci (var) atau menggunakan deklarasi singkat :=.

Tipe data menentukan jenis nilai yang dapat disimpan oleh sebuah variabel. Beberapa tipe data dasar yang umum digunakan dalam Go adalah:

int → menyimpan bilangan bulat.
float64 → menyimpan bilangan desimal.
string → menyimpan teks atau kumpulan karakter.
bool → menyimpan nilai logika true atau false.

Operasi merupakan proses yang dilakukan terhadap nilai atau variabel. Dalam Go terdapat beberapa jenis operator, seperti operator aritmatika, perbandingan, dan logika.

Operator aritmatika antara lain:

+ untuk penjumlahan
- untuk pengurangan
* untuk perkalian
/ untuk pembagian
% untuk mencari sisa pembagian


### B. [Judul Topik Dasar Teori 2, misal: Package dan Struktur Program di Go]
Dalam kerangka program yang ditulis dalam bahasa pemrograman Go, program utama selalu mempunyai dua komponen berikut:
 • package main merupakan penanda bahwa file ini berisi program utama.
 • func main() berisi kode utama dari sebuah program Go.

Komentar, bukan bagian dari kode program, dan dapat ditulis di mana saja di dalam program: • Satu baris teks yang diawali dengan garis miring ganda ('//') s.d. akhir baris, atau. • Beberapa baris teks yang dimulai dengan pasangan karakter '/*' dan diakhiri dengan '*/'. 

#### 1. [Judul Sub-topik 1, misal: Pengertian Package main dan func main()]
package main adalah deklarasi yang menunjukkan bahwa sebuah file Go termasuk dalam package utama
func main() adalah fungsi utama dalam program Go, sehingga perintah akan dijalankan ketika program dimulai.

#### 2. [Judul Sub-topik 2, misal: Tipe Data dan Deklarasi Variabel di Go]
Tipe data adalah jenis data yang menentukan nilai yang dapat disimpan oleh sebuah variabel.
Deklarasi variabel adalah proses membuat variabel untuk menyimpan suatu nilai

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. [Skor.go]

```go
package main

import "fmt"

func main() {
	var nama string
	var matematika int
	var bahasaInggris int

	fmt.Scanln(&nama)
	fmt.Scanln(&matematika)
	fmt.Scanln(&bahasaInggris)

	Total := matematika + bahasaInggris
	rataRata := Total / 2

	fmt.Println(nama)
	fmt.Println(Total)
	fmt.Println(rataRata)
}
```
#### Deskripsi
[Tujuan dari program tersebut adalah menerima input data dari mahasiswa dan nilai ujiannya, lalu menghitung rata-rata nilai tersebut.
Mengimplementasikan ekspresi logika sekuensial untuk menjumlahkan kedua nilai mata kuliah (bhs inggris dan matematika) dalam variabel total lalu membaginya dengan angka 2 untuk mencari rata-rata.
Hasil akhir (output) berupa nama mahasiswa, total nilai dan rata-rata.


### 2. [Nama File, misal: Tukar.go]

```go
[package main

import "fmt"

func main() {
	var a, b int

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)
}
]
```
#### Deskripsi
[program dibuat untuk menerima dua inputan bilangan bulat, yaitu variabel a dab b. untuk menukar variabel a,b ke b,a
impementasi program yang dibuat adalah membaca dua nilai integer, menukar kedua variabel, lalu mencetaknya kembali
hasil yang diperoleh adalah hasil tukar dari kedua variabel tersebut]

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. [Cacahuang]

```go
[package main
import "fmt"

func main() {
	var uang int

	fmt.Scan(&uang)

	sepuluhRibu := uang / 10000
	sisa := uang % 10000

	limaRibu := sisa / 5000
	sisa = sisa % 5000

	seribu := sisa / 1000

	fmt.Println(sepuluhRibu, limaRibu, seribu)
}]
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](https://github.com/renwxyz/alpro-11-01/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/cacahuang/output/cacahuang.go.png)


#### Deskripsi
[program dibuat untuk mensimulasikan pemecahan atau konversi jumlah uang yaitu pecahan sepuluh ribu, lima ribu, seribu
Program meminta input nilai uang dari pengguna yang disimpan dalam variabel. lalu gunakan operasi pembagian (/) uang untuk menghitung jumlah pecahan lembaran, serta operasi modulo (%) untuk mencari nilai sisa uang yang belum terkonversi. langkah-langkah ini dapat diulang dari jumlah lembaran terbesar hingga terkecil
Program berhasil memecah nominal uang total menjadi lembaran pecahan sepuluh ribuan, lima ribuan, dan seribuan secara berurutan]

### 2. [Kalkulator]

```go
[package main
import "fmt"

func main() {
	var a,b int

	fmt.Scan (&a, &b)

	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a*b)
	fmt.Println(a/b)
	fmt.Println(a%b)
}]
```

##### Output
![Screenshot Output Unguided](https://github.com/renwxyz/alpro-11-01/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/cacahuang/output/kalkulator.go.png)

#### Deskripsi
[program dibuat untuk menerima dua inputan bilangan bulat, yaitu variabel a dab b. lalu mencetak hasilnya menggunakan fmt.Print.ln.
impementasi fungsi kalkulator dasar yang menggunakan operasi penjumlahan, pengurangan, perkalian, pembagian, dan sisa bagi
hasilnya program dapat membaca dua input dan berhasil memecahkan kelima operasi hitung yang terstruktur]

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
[Tuliskan kesimpulan yang menjawab tujuan praktikum berdasarkan hasil yang diperoleh.]

## Referensi
1. [Yohanes Calvin]. ([2026]). *[Judul Buku/Sumber]*. [Kota]: [Penerbit]. Diakses pada [tanggal akses] melalui [tautan/DOI]
2. [Yohanes Calvin]. ([2026]). *[Judul Buku/Sumber]*. [Kota]: [Penerbit]. Diakses pada [tanggal akses] melalui [tautan/DOI]
<!-- Tambahkan nomor referensi berikutnya sesuai kebutuhan -->
