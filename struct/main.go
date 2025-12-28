package main

import "fmt"

type User struct {
	FirstName string
	LastName string
	Age int
}

func (user User) FullName(title string) string {
	return user.FirstName + " " + user.LastName + " " + title
}

func main() {
	user := User{
		FirstName: "Alvin",
		LastName: "Sanudharma",
		Age: 27,
	}

	fmt.Println("Nama Depan :", user.FirstName)
	fmt.Println("Nama Belakang :", user.LastName)
	fmt.Println("Nama Lengkap :", user.FullName("S.Kom"))
	fmt.Println("Umur :", user.Age)
}