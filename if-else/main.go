package main

import "fmt"

func main()  {
	var nilai int

	fmt.Print("Masukan nilai ujian: ")
	fmt.Scanln(&nilai)

	if nilai > 90 {
		fmt.Println("Nilai sangat baik!")
	} else {
		fmt.Println("Nilai kurang ☹️")
	}
}