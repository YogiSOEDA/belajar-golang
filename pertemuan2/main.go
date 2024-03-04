package main

import "fmt"

func main() {
	urutanAngka := map[string]int{
		"satu": 1,
		"dua": 2,
		"tiga": 3,
		"empat": 4,
	}

	for key, value := range urutanAngka{
		fmt.Println("key: ",key)
		fmt.Println("value: ",value)
	}
}