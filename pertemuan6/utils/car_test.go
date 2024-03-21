package utils_test

import (
	"context"
	"fmt"
	"sekolahbeta/pertemuan6/config"
	"sekolahbeta/pertemuan6/model"
	"sekolahbeta/pertemuan6/utils"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init()  {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("env not found, using system env")
	}
}

func TestCreateDataSuccess(t *testing.T) {
	Init()

	conn, err := config.OpenConn()

	assert.Nil(t, err)

	bdy := model.Car{
		ID: "1",
		Nama: "toyota",
		Tipe: "yaris",
		Tahun: "2018",
	}

	err = utils.InsertData(conn, bdy, context.TODO())
	assert.Nil(t, err)

}

func TestCreateDataFailed(t *testing.T)  {
	Init()

	conn, err := config.OpenConn()
	assert.Nil(t, err)

	bdy := model.Car{
		ID: "123",
		Nama: "toyota",
		Tipe: "yaris",
		Tahun: "2018",
	}

	err = utils.InsertData(conn, bdy, context.TODO())
	assert.Nil(t, err)

	bdy1 := model.Car{
		ID: "123",
		Nama: "toyota",
		Tipe: "yaris",
		Tahun: "2018",
	}

	err1 := utils.InsertData(conn, bdy1, context.TODO())
	assert.NotNil(t, err1)
}

func TestGetByID(t *testing.T)  {
	Init()

	conn, err := config.OpenConn()
	assert.Nil(t, err)

	id := "123"

	res, err := utils.GetByID(conn, id, context.TODO())
	assert.Nil(t, err)
	assert.NotEqual(t, model.Car{}, res)
	fmt.Println(res)
}