package main

import (
	"fiber-mongo-api/configs"
	"fiber-mongo-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	app.Use(cors.New())

	//routes
	routes.UserRoute(app)
	routes.PageRoute(app)
	routes.ConsultancyRoute(app)

	app.Listen(":8080")
}
