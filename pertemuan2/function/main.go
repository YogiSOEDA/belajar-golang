package main

import "fmt"

func hello() {
	fmt.Println("Awoooo")
}

func konversiMataUangUSD(uang int) int  {
	var hasil int

	hasil = uang * 15000

	return hasil
}

//implementasi fungsi 2 input 2 output
func konversiMataUang(uang int, currency string) (int, string) {
	var hasil int
	
	switch currency{
	case "USD":
		hasil = uang * 15000
	case "JPY":
		hasil = uang * 300
	default:
		hasil = 0
	}

	return hasil, currency
}

//implementasi variadic function
func kalkulator(operator string, angka ...int) int {
	var hasil int
	for i := 0; i < len(angka); i++{
		if operator == "+" {
			hasil += angka[i]
		} else if operator == "-"{
			hasil -= angka[i]
		}
	}

	return hasil
}

func main() {
	hello()

	uang := 2000
	
	hasilKonversi := konversiMataUangUSD(5)
	fmt.Println("Hasil konversi= ",hasilKonversi)

	hasilKonversi1, currency := konversiMataUang(uang, "JPY")
	fmt.Println(hasilKonversi1)
	fmt.Println(currency)

	angkaInput := []int{1,2,3,4,5,6,7,8,9}

	fmt.Println(kalkulator("+", angkaInput...))

	// fmt.Println(konversiMataUangUSD(2))
}
