package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

// import (
// 	"fmt"
// 	"time"
// )

//contoh goroutine
// func main() {
// 	startedAt := time.Now()
// 	for i := 0; i < 15; i++ {
// 		go uploadFile(i)
// 	}
// 	// time.Sleep(1 * time.Second)
// 	fmt.Println(time.Since(startedAt))

// }

// func uploadFile(i int) {
// 	fmt.Printf("Uploading file %d\n", i)
// 	time.Sleep(500 * time.Millisecond)
// }

type Car struct {
	ID           string `json:"id"`
	Year         string `json:"year"`
	Make         string `json:"make"`
	Model        string `json:"model"`
	Trim         string `json:"trim"`
	Body         string `json:"body"`
	Transmission string `json:"transmission"`
	State        string `json:"state"`
	Condition    string `json:"condition"`
	Odometer     string `json:"odometer"`
	Color        string `json:"color"`
	Interior     string `json:"interior"`
	Seller       string `json:"seller"`
	Mmr          string `json:"mmr"`
	SellingPrice string `json:"selling_price"`
	SaleDate     string `json:"sale_date"`
}

func main() {
	
	err := os.RemoveAll("json")
	if err != nil {
		fmt.Println(err)
	}

	err = os.Mkdir("json", 0777)
	if err != nil {
		fmt.Println(err)
	}

	fileCsv, err := os.Open("cars_100000.csv")
	if err != nil {
		fmt.Println(err)
	}

	defer fileCsv.Close()

	reader := csv.NewReader(fileCsv)
	
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	
	
	// fmt.Println(records[:2])
	cars := csvToStruct(records)
	// fmt.Println(cars[:2])
	
	startedAt := time.Now()
	// fmt.Println(string(convertToJson(cars[1])))
	
	// encoded := convertToJson(cars[1])
	
	// saveJsonToFile(encoded, cars[1].ID)
	
	//menunggu goroutine
	wg := sync.WaitGroup{}

	channel := make(chan Car)

	wg.Add(3)

	go process(channel, &wg)
	go process(channel, &wg)
	go process(channel, &wg)
	// go process(channel, &wg)
	// go process(channel, &wg)
	// go process(channel, &wg)

	for _, car := range cars {
		// wg.Add(1)

		//goroutine cara 1
		// go saveJsonToFile(convertToJson(car), car.ID)
		
		//goroutine cara 2 dengan anonymous function
		// go func (car Car, wg *sync.WaitGroup)  {
		// 	encoded := convertToJson(car)
		// 	saveJsonToFile(encoded, car.ID)

		// 	wg.Done()
		// }(car, &wg)

		channel <- car
	}

	close(channel)

	wg.Wait()

	fmt.Println("Sukses")
	fmt.Println(time.Since(startedAt))

	// fileCsv.Close()
}

func csvToStruct(records [][]string) []Car {
	cars := []Car{}

	for _, car := range records {
		cars = append(cars, Car{
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

func convertToJson(car Car) []byte {
	encoded, err := json.MarshalIndent(car, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	return encoded
}

func saveJsonToFile(encoded []byte, name string) {
	file, err := os.Create(fmt.Sprintf("json/%s.json", name))
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	_, err = file.Write(encoded)
	if err != nil {
		fmt.Println(err)
	}

}

func process(channel chan Car, wg *sync.WaitGroup)  {
	for data := range channel {
		encoded := convertToJson(data)
		saveJsonToFile(encoded, data.ID)
	}

	wg.Done()
}
