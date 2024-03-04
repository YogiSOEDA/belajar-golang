package main

import "fmt"

type Person struct {
	Name      string
	DateBirth string
	Weight    string
	Mail []Email
}

type Email struct{
	To string
	Name string
}

func main() {
	person1 := Person{
		Name:      "Ucup",
		DateBirth: "01-01-1998",
		Weight:    "60",
		Mail: []Email{
			{
				To: "ucup@mail.com",
				Name: "ucup",
			},
			{
				To: "asep@mail.com",
				Name: "ucup",
			},
		},
	}

	fmt.Println(person1.Mail[1].To)

	persons := []Person{
		{Name: "Asep", DateBirth: "01-01-1990", Weight: "50"},
		{Name: "Asep", DateBirth: "01-01-1990", Weight: "50"},
		{Name: "Asep", DateBirth: "01-01-1990", Weight: "50"},
		{Name: "Asep", DateBirth: "01-01-1990", Weight: "50"},
	}

	fmt.Println(persons)
}