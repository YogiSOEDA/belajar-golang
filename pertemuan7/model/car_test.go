package model_test

import (
	"fmt"
	"sekolahbeta/hacker/database-orm/config"
	"sekolahbeta/hacker/database-orm/model"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("env not found, using global env")
	}
	config.OpenDB()
}

func TestCreateCar(t *testing.T) {
	Init()

	carData := model.Car{
		Nama:  "toyota",
		Tipe:  "crown",
		Tahun: "1998",
	}

	err := carData.Create(config.Mysql.DB)
	assert.Nil(t, err)
}

func TestGetByIdCar(t *testing.T) {
	Init()

	carData := model.Car{
		Model: model.Model{
			ID: 1,
		},
	}

	data, err := carData.GetByID(config.Mysql.DB)
	assert.Nil(t, err)

	fmt.Println(data)
}

func TestGetAll(t *testing.T) {
	Init()

	carData := model.Car{
		Nama:  "toyota",
		Tipe:  "supra",
		Tahun: "1999",
	}

	err := carData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	res, err := carData.GetAll(config.Mysql.DB)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(res), 1)

	fmt.Println(res)
}

func TestUpdate(t *testing.T) {
	Init()

	carData := model.Car{
		Nama:  "honda",
		Tipe:  "asoy",
		Tahun: "2000",
	}

	err := carData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	fmt.Println(carData.ID)

	carData.Nama = "wak waw"
	// carData = model.Car{
	// 	Model: model.Model{
	// 		ID: carData.ID,
	// 	},
	// 	Nama: "toyota",
	// 	Tipe: "paling baru",
	// 	Tahun: "2020",
	// }

	err = carData.UpdateOneByID(config.Mysql.DB)
	assert.Nil(t, err)
}

func TestDeleteByID(t *testing.T)  {
	Init()

	carData := model.Car{
		Nama:  "honda",
		Tipe:  "asoy",
		Tahun: "2000",
	}
	err := carData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	fmt.Println(carData.ID)

	// carDataDelete := model.Car{
	// 	Model: model.Model{
	// 		ID: carData.ID,
	// 	},
	// }
	err = carData.DeleteByID(config.Mysql.DB)
	assert.Nil(t, err)
}