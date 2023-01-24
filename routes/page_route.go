package routes

import (
	"fiber-mongo-api/controllers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func PageRoute(app *fiber.App) {
	fmt.Println("res")
	fmt.Println("newpage")

	//app.Get("/page/:pageId", controllers.GetApage)
	app.Get("/pages", controllers.GetAllPages)
	app.Post("/pages", controllers.CreatePages)
}
