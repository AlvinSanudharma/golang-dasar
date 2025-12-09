package main

import "fmt"

func main()  {
	mahasiswa := map[string]string{
		"nama": "Alvin",
		"jurusan": "Sistem Informasi",
	}

	fmt.Println("Nama :", mahasiswa["nama"])
	fmt.Println("Jurusan :", mahasiswa["jurusan"])

	mahasiswa["ukm"] = "SBMC"

	fmt.Println("UKM :", mahasiswa["ukm"])
}