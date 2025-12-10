package main

import "fmt"

func main()  {
	x := 10

	x = 5
	fmt.Println(x)

	x += 2 // x = x + 2
	fmt.Println("x += 2 :", x)

	x -= 2 // x = x - 2
	fmt.Println("x -= 2 :", x)
}