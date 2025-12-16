package main

import "fmt"

func tambah(a, b int) (hasil int) {
	hasil = a + b

	return 
}

func sayHello(name string) (fullName string) {
	fullName = "Hello my name is " + name

	return 
}

func sum(angka ...int) int {
	total := 0

	for _,v := range angka {
		total += v
	}

	return total
}

func main(){
	fmt.Println("Hasil penjumlahan (Variadic) : ", sum(1, 2, 3, 4, 5))
	fmt.Println("Hasil penjumlahan : ", tambah(1, 2))
	fmt.Println(sayHello("Alvin"))
}