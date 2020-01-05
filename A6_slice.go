package main
import "fmt"
func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}
	// var fruitsA = []string{"apple", "grape"}      // slice
	// var fruitsB = [2]string{"banana", "melon"}    // array
	// var fruitsC = [...]string{"papaya", "grape"}  // array
	fmt.Println(fruits[0]) // "apple"
	var newFruit = fruits[0:2]
	fmt.Println(fruits[:]) //semua element
	fmt.Println(newFruit) //["apple", "grape"]
	fmt.Println(fruits[2:]) //semua element mulai dari index 2
	fmt.Println(fruits[:2]) //semua element sebelum index 2


	// Slice merupakan tipe data reference atau referensi. Artinya jika ada slice baru yang terbentuk dari slice lama, 
	// maka data elemen slice yang baru akan memiliki alamat memori yang sama dengan elemen slice lama. 
	// Setiap perubahan yang terjadi di elemen slice baru, akan berdampak juga pada elemen slice lama yang memiliki referensi yang sama.
	
	// append
	var buah = append(fruits,"jeruk");
	fmt.Println(fruits, buah)

	// cap() digunakan untuk menghitung lebar atau kapasitas maksimum slice.
	// copy(dst, src)
}