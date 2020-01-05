package main
import "fmt"

func main() {
	var array1 [2]string // mengisi array dengan 2 slot
	array1[0] = "satu"
	array1[1] = "dua"
	// array1[2] = "3"  akan error karena deklarasi 2 slot

	fmt.Println(array1[0], array1[1])

	var array2 = [2]string{"1","2"}
	var array3 = [...]string{"1","2"} // inisialisasi tanpa jumlah

	fmt.Println(array2[0], array2[1])
	fmt.Println("Jumlah element \t\t", len(array2))
	fmt.Println("Isi semua element \t", array2)
	fmt.Println("Isi semua element array3 \t", array3)

	// array multi dimensi

	var array4 = [2][3]string{[3]string{"irfangi", "said", "saidi"},[3]string{"22","223", "22"}} // menyesuaikan dimensi arraynya
	fmt.Println("Isi semua element array3 \t", array4)

	// menampilkan
	// menggunakan for 

	for i :=0 ; i< len(array3); i++{
		fmt.Println("perulangan array ", array3[i])
	}

	// menggunakan range

	for i, isi3 := range array3 {
		fmt.Println("perulangan array range ", i , isi3)
	}
} 