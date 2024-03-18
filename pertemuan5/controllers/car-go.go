package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sekolahbeta/hacker/model"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

func openFile(path string) (*os.File, error)  {
	return os.Open(path)
}

func loadFileGoroutine(file *os.File) (<-chan []string, error)  {
	carChan := make(chan []string)
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return carChan, err
	}

	go func() {
		for _, car := range records {
			carChan <- car
		}

		close(carChan)
	}()

	return carChan, nil
}

func processConvertStruct(csvChan <- chan []string) <- chan model.Car  {
	carsChan := make(chan model.Car)

	go func() {
		for car := range csvChan {
			carsChan <- model.Car{
				ID: car[0],
				Year: car[1],
				Make: car[2],
				Model: car[3],
				Trim: car[4],
				Body: car[5],
				Transmission: car[6],
				State: car[7],
				Condition: car[8],
				Odometer: car[9],
				Color: car[10],
				Interior: car[11],
				Seller: car[12],
				Mmr: car[13],
				SellingPrice: car[14],
				SaleDate: car[15],
			}
		}

		close(carsChan)
	}()

	return carsChan
}

func appendCars(carChanMany ...<-chan model.Car) <-chan model.Car  {
	wg := sync.WaitGroup{}

	mergedChan := make(chan model.Car)

	wg.Add(len(carChanMany))
	for _, ch := range carChanMany {
		go func(ch <- chan model.Car) {
			for cars := range ch {
				mergedChan <- cars
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(mergedChan)
	}()

	return mergedChan
}

func convertCarsToJson(data []model.Car) ([]byte, error)  {
	return json.Marshal(data)
}

func GetCarsGo(c *fiber.Ctx) error {
	started := time.Now()
	file, err := openFile("resources/cars_500.csv")
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer file.Close()

	csvChan, err := loadFileGoroutine(file)
	if err != nil {
		return err
	}

	jmlGoroutine := 10

	var carChanTemp []<-chan model.Car

	for i := 0; i < jmlGoroutine; i++ {
		carChanTemp = append(carChanTemp, processConvertStruct(csvChan))
	}

	mergedCh := appendCars(carChanTemp...)

	var cars []model.Car

	for ch := range mergedCh {
		cars = append(cars, ch)
	}

	jsonData, err := convertCarsToJson(cars)
	if err != nil {
		return nil
	}

	fmt.Println("[Dengan Goroutine]", time.Since(started))

	return c.Send(jsonData)
}