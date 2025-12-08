package main

import "fmt"

func main() {
	var angka [5]int = [5]int{10, 20, 30, 40, 50}

	fmt.Println(angka)
	fmt.Println(angka[1])
	fmt.Println("Jumlah elemen : ",len(angka))
}