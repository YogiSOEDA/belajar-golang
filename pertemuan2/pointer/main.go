package main

import "fmt"

func main() {
	var nama string = "asep"
	var namaPointer *string = &nama

	fmt.Println("ini nama dengan &: ", &nama)
	fmt.Println("ini variabel namaPointer: ", namaPointer)
	fmt.Println(nama)

	*namaPointer = "ucup"

	fmt.Println(*namaPointer)
	fmt.Println(namaPointer)
	fmt.Println(nama)

	nama = "udin"
	
	fmt.Println(&nama)
	fmt.Println(*namaPointer)
	fmt.Println(nama)
	fmt.Println(namaPointer)
}