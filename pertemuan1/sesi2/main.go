package main

import "fmt"

func main() {
	mataUang := map[string]string{
		"USD" : "$",
		"IDR" : "Rp. ",
	}

	fmt.Println(mataUang)

	//menambah key map baru
	mataUang = map[string]string{
		"JPY" : "Y",
		"USD" : "Dollar",
	}

	fmt.Println(mataUang)
	fmt.Println(mataUang["USD"])

	delete(mataUang, "uwu")
	fmt.Println(mataUang)

	value, isExist := mataUang["USD"]
	fmt.Println(value, isExist)

	//kombinasi slice map
	testing1 := []map[string]int{
		{"testing1":2376},
		{
			"testing1":2376,
			"testing2":7922,
		},
		{"testing1":2376},
		{"testing1":2376},
		{"testing1":2376},
		{"testing1":2376},
		{"testing1":2376},
	}

	fmt.Println(testing1)
	fmt.Println(testing1[1]["testing2"])

	//operator aritmatika, operator perbandingan, operator logika
	//percabangan if dan switch case
	//perulangan for


	
}