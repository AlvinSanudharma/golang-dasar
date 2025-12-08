package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Halo, Golang!"

	fmt.Println("Lowercase:", strings.ToLower(text))
	fmt.Println("Uppercase:", strings.ToUpper(text))
	fmt.Println("Replace:", strings.ReplaceAll(text, "Halo", "Hi"))
}
