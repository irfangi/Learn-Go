package main

import "fmt"

func main() {
	// for
	fmt.Println("for")
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}
	// basic while
	fmt.Println("while")

	var i = 0;

	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}

	// basic do while
	var item = 0;
	fmt.Println("do while")
	for{
		fmt.Println("Angka", item)
		item++
		if item > 4 {
			break
		}
		//lawan dari break adalah continue
	}
}