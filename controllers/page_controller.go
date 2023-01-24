package controllers

import (
	"context"
	"fiber-mongo-api/configs"
	"fiber-mongo-api/models"
	"fiber-mongo-api/responses"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var pageCollection *mongo.Collection = configs.GetCollection(configs.DB, "pages")

var validatePageModel = validator.New()

func GetAllPages(c *fiber.Ctx) error {
	fmt.Print("pages")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var pages []models.Pages

	defer cancel()

	results, err := pageCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.GetResponse{
			Status:  http.StatusInternalServerError,
			Message: "error", Data: err.Error()})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singlepage models.Pages
		if err = results.Decode(&singlepage); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				responses.GetResponse{Status: http.StatusInternalServerError,
					Message: "error", Data: err.Error()})
		}

		pages = append(pages, singlepage)
	}

	return c.Status(http.StatusOK).JSON(
		responses.GetResponse{Status: http.StatusOK, Message: "success",
			Data: pages},
	)
}

func CreatePages(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// var page models.Page

	var page1 []models.Pages
	defer cancel()
	if err := c.BodyParser(&page1); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.PageResponse{Status: http.StatusBadRequest,
			Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	for _, p := range page1 {
		err := validate.Struct(p)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(responses.PageResponse{Status: http.StatusBadRequest,
				Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	docs := make([]interface{}, len(page1))
	for i, b := range page1 {
		docs[i] = bson.M{

			"Name":     b.Name,
			"Pid":      b.Pid,
			"Ptitle":   b.Ptitle,
			"Pdesc":    b.Pdesc,
			"Pkeyword": b.Pkeyword,
			"Pimg":     b.Pimg,
			"PimgAlt":  b.PimgAlt,
			"Pauthor":  b.Pauthor,
		}
	}

	result, err := pageCollection.InsertMany(ctx, docs)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.PageResponse{Status: http.StatusInternalServerError,
			Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}

	return c.Status(http.StatusCreated).JSON(responses.PageResponse{Status: http.StatusCreated,
		Message: "success", Data: &fiber.Map{"data": result}})
}
