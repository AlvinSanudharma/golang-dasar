package main

import "fmt"

func main(){
	angka := 1

	switch angka {
		case 1:
			fmt.Println("Ini angka 1")
		case 2:
			fmt.Println("Ini angka 2")
		case 3:
			fmt.Println("Ini angka 3")
		case 4:
			fmt.Println("Ini angka 4")
		default:
			fmt.Println("Angka tidak ditemukan!")
	}
}