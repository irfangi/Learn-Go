// Tipe data	Cakupan bilangan
// uint8	0 ↔ 255
// uint16	0 ↔ 65535
// uint32	0 ↔ 4294967295
// uint64	0 ↔ 18446744073709551615
// uint	sama dengan uint32 atau uint64 (tergantung nilai)
// byte	sama dengan uint8
// int8	-128 ↔ 127
// int16	-32768 ↔ 32767
// int32	-2147483648 ↔ 2147483647
// int64	-9223372036854775808 ↔ 9223372036854775807
// int	sama dengan int32 atau int64 (tergantung nilai)
// rune	sama dengan int32

package main

import "fmt"

func main() {
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644 // jika tidak tipe data tak di deklarasikan maka akan menyesuaikan ke int32 jika di kasih uint8 maka error

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	var decimalNumber = 2.62 //float 32

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber) // f untuk memformat float

	var exist bool = true
	fmt.Printf("exist? %t \n", exist) // t untuk memfirmat bolean

	var message string = "Halo"
	fmt.Printf("message: %s \n", message)

	// zero value
	// string ""
	// bolean false
	// numeric 0 
	// decimal 0.0

	// nill => 
		// pointer
		// tipe data fungsi
		// slice
		// map
		// channel
		// interface kosong atau interface{}
	// data nill sudah tidak bisa di isi dengan variabel sebelumya

	// constanta digunakan untuk variabel yang tidak di ubah ubah contoh pi

	
}

