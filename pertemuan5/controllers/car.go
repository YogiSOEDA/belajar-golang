package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sekolahbeta/hacker/model"
	

	"github.com/gofiber/fiber/v2"
)

func GetCars(c *fiber.Ctx) error {
	fileCsv, err := os.Open("resources/cars_500.csv")
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer fileCsv.Close()

	reader := csv.NewReader(fileCsv)

	// ch := make(chan csv.Reader)

	// wg := sync.WaitGroup{}

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return err
	}

	cars := csvToStruct(records)

	// encoded := convertToJson(cars[])

	encoded, err := json.MarshalIndent(cars, "", "    ")
	if err != nil {
		fmt.Println(err)
		return err
	}

	// for _, car := range cars {
	// 	encoded := convertToJson(car)
	// }

	return c.Send(encoded)
}

// func loadFromCsv(ch <-chan csv.Reader, wg *sync.WaitGroup) {
// 	for reader := range ch {
// 		records, err := reader.ReadAll()
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		csvToStruct(records)
// 	}

// 	wg.Done()
// }

func csvToStruct(records [][]string) []model.Car {
	cars := []model.Car{}

	for _, car := range records {
		cars = append(cars, model.Car{
			ID:           car[0],
			Year:         car[1],
			Make:         car[2],
			Model:        car[3],
			Trim:         car[4],
			Body:         car[5],
			Transmission: car[6],
			State:        car[7],
			Condition:    car[8],
			Odometer:     car[9],
			Color:        car[10],
			Interior:     car[11],
			Seller:       car[12],
			Mmr:          car[13],
			SellingPrice: car[14],
			SaleDate:     car[15],
		})
	}

	return cars
}

// func convertToJson(car model.Car) []byte {
// 	encoded, err := json.MarshalIndent(car, "", "    ")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return encoded
// }
