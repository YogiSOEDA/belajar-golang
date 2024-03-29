package main

import (
	"fmt"
	fiberConfig "sekolahbeta/hacker/config/fiber"
	"sekolahbeta/hacker/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Get("/pesanan", controllers.GetPesanan)

	app.Get("/cars", controllers.GetCars)

	app.Get("/cars-go", controllers.GetCarsGo)

	listenAddress := fmt.Sprintf("%s:%s", fiberConfig.GetFiberHttpHost(), fiberConfig.GetFiberHttpPort())

	err := app.Listen(listenAddress)
	if err != nil {
		panic(err)
	}
}