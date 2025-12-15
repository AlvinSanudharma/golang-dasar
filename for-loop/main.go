package main

import "fmt"

func main(){
	for i := 0; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	buah := []string{"apel", "jeruk", "pisang"}

	for index, value := range buah {
		fmt.Printf("Index %d: %s\n", index, value)
	}
}