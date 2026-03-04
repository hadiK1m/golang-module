package main

import "fmt"

func main() {
	// Deklarasi dan inisialisasi variabel
	var warna [4]string
	warna[0] = "merah"
	warna[1] = "kuning"
	warna[2] = "hijau"
	warna[3] = "biru"

	// Menampilkan nilai variabel
	fmt.Println("Array warna:", warna)
	fmt.Println("Warna ke 2:", warna[2])

	// slice (potongan) adalah bagian dari array yang memiliki ukuran dinamis
	buah := []string {"Apel", "Mangga", "Pisang", "Jeruk"}
	fmt.Println("Slice buah:", buah)

	// Menambahkan elemen ke slice
	buah = append(buah, "anggur", "nanas")
	fmt.Println("Slice buah:", buah)

	// Mengambil slice dari array
	favorite := buah[0:4]
	fmt.Println("Buah favorit:", favorite)

	// Panjang dan kapasitas slice
	fmt.Println("Jumlah buah:", len(buah))
	fmt.Println("Capacity:", cap(buah))

	// Perulangan dengan range	
	for index, item := range buah {
		fmt.Printf("Index %d: %s\n", index, item)
	}
}