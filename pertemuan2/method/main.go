package main

import "fmt"

type hewan struct {
	Nama  string
	Suara string
}

func (h hewan) bunyi() {
	fmt.Println(h.Suara)
}

func main() {
	hwn := hewan{
		Nama: "Kucing",
		Suara: "Meow",
	}

	hwn.bunyi()
}