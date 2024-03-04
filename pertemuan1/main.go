package main

import "fmt"

func main() {

	var paragraf string = `
	ini adalah
	contoh paragraf
	`

	// var angka int8 = 5
	// angka = 5

	//deklarasi float tanpa inisiasi tipe data
	// angka1 := -89.7
	// var angka1 float32 = -89.7

	// var belajar bool = true

	//cara lain deklarasi variabel
	var(
	angka int8 = 5
	angka1 float32 = -89.7
	belajar bool = true
	angkaArray [2]int
	)

	angkaArray[0] = 1
	angkaArray1 := [2]int{2,5}
	angkaSlice1 := []string{"saya", "tamvan", "1", "3333"}//ini slice
	hurufSlice1 := []string{"ini", "tambahan", "slice"}
	angkaSlice2 := make([]string,4)

	nama := "yogi"

	fmt.Println(paragraf)
	fmt.Println(angka)
	fmt.Printf(`
	nama saya: %s, 
	angka saya adalah: %d %f, 
	saya belajar golang: %v 
	`, nama, angka, angka1, belajar)
	fmt.Println("testing", angka, " ini ada angka")
	fmt.Println("satu")


	//belajar array
	fmt.Println(angkaArray[0])
	fmt.Println(angkaArray1)
	fmt.Println(angkaSlice1)

	//ini slice
	fmt.Println("panjang slice adalah: ", len(angkaSlice1))

	angkaSlice1 = append(angkaSlice1, hurufSlice1...)
	fmt.Println(angkaSlice1)
	
	copy(angkaSlice2, angkaSlice1)
	fmt.Println(angkaSlice2)


	//belajar map
	
}