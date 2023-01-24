package routes

import (
	"fiber-mongo-api/controllers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ConsultancyRoute(app *fiber.App) {
	fmt.Println("res")
	fmt.Println("newJobseeker")

	//app.Get("/page/:pageId", controllers.GetApage)
	// app.Get("/pages", controllers.GetAllPages)
	app.Post("/jobSeeker", controllers.CreateJobSeeker)
	app.Post("/hire", controllers.CreateHire)

}
