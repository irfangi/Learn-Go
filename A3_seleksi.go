package main

import "fmt"

func main() {
	var nilai = 7
	
	if nilai == 10{
		fmt.Println("lulus sempurna")
	}else if nilai >= 5 {
		fmt.Println("lulus")
	}else if nilai == 4{
		fmt.Println("hampir lulus")
	}else{
		fmt.Printf("tidak lulus niali anda %d\n", nilai)
	}

	// tmp variabel if

	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// switch

	switch nilai {
	//banyak kondisi
	case 8,9,10:
		fmt.Println("good")
	case 7:
		//kurung kurawal 
		{
			fmt.Println("nice")
			fmt.Println("nice")
		}
		
	default:
		fmt.Println("not bad")
	}
}