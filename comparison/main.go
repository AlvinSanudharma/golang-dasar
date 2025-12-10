package main

import "fmt"

func main()  {
	var angka1, angka2 int

	fmt.Print("Masukan angka pertama: ")
	fmt.Scanln(&angka1)
	fmt.Print("Masukan angka kedua: ")
	fmt.Scanln(&angka2)
	fmt.Println("\n=== Hasil Perbandingan ===")
	fmt.Printf("%d == %d ? %v\n", angka1, angka1, angka1 == angka2)
	fmt.Printf("%d != %d ? %v\n", angka1, angka1, angka1 != angka2)
	fmt.Printf("%d > %d ? %v\n", angka1, angka1, angka1 > angka2)
	fmt.Printf("%d < %d ? %v\n", angka1, angka1, angka1 < angka2)
	fmt.Printf("%d >= %d ? %v\n", angka1, angka1, angka1 >= angka2)
	fmt.Printf("%d <= %d ? %v\n", angka1, angka1, angka1 >= angka2)
}