package main

import "fmt"

func main(){
	defer fmt.Println("defer ini tetap jalan sebelum program mati")
	fmt.Println("Sebelum panic")
	panic("Ada sesuatu yang fatal!")
	// baris dibawah tidak dieksekusi
	fmt.Println("Setelah panic")
}