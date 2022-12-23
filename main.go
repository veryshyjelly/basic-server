package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/changeCity", cityHandler)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://192.168.140.99:19000",
	}))

	log.Fatalln(app.Listen(":3000"))
}

func cityHandler(ctx *fiber.Ctx) error {
	fmt.Println(ctx.Request().Body())
	return nil
}
