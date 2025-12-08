package main

import "fmt"

func main()  {
	arr := [6]int{1,2,3,4,5,6}

	s1 := arr[:] // ambil seluruh element

	fmt.Println("S1 :",s1)

	s2 := arr[:3] // ambil element pertama sampai ketiga

	fmt.Println("S2 :",s2)
}