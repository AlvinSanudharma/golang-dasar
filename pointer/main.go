package main

import "fmt"

func ubahNama(nama *string)  {
	*nama = "Budi"
}

func main(){
	nama := "Andi"

	ubahNama(&nama)
	fmt.Println(nama)
}