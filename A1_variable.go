package main

import "fmt"

func main() {
	var frstName = "irfangi"
	var lastName string
	
	lastName = "fangi"

	tanpaVar := "fangi"


	// Fungsi ini digunakan untuk menampilkan output dalam bentuk tertentu. Kegunaannya sama seperti fungsi fmt.Println(), hanya saja struktur outputnya didefinisikan di awal.

	// Perhatikan bagian "halo %s %s!\n", karakter %s disitu akan diganti dengan data string yang berada di parameter ke-2, ke-3, dan seterusnya.

	// Ketiga baris kode di bawah ini menghasilkan output yang sama, meskipun cara penulisannya berbeda.

	// fmt.Printf("halo john wick!\n")
	// fmt.Printf("halo %s %s!\n", firstName, lastName)
	// fmt.Println("halo", firstName, lastName + "!")

	// fmt.Printf("Hallo %s %s!\n", frstName, lastName)
	fmt.Printf("Hallo %s %s %s!\n", frstName,lastName ,tanpaVar)


	// deklarasi multi variabel

	// var satu, dua, tiga string = "satu", "dua" , "tiga"	
	
	// var satu, dua, tiga string
	// satu, dua, tiga = "satu", "dua" , "tiga"	
	
	satu, dua, tiga := "satu", "dua" , "tiga"	
	fmt.Println(satu,dua,tiga)

	// variabel sampah _

	var _ string = "sampah"
	not_trash :="2"
	fmt.Println(not_trash)

	// deklarasi menggunakan new
	baru := new(string)

	fmt.Println(baru) //return 0xc42000e250 alamat memeori
	fmt.Println(*baru) //return ""
}
