package controllers

import (
	"context"
	"fiber-mongo-api/configs"
	"fiber-mongo-api/models"
	"fiber-mongo-api/responses"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var JobSeekerCollection *mongo.Collection = configs.GetCollection(configs.DB, "JobSeeker")

var HireCollection *mongo.Collection = configs.GetCollection(configs.DB, "Hire")

var validateJobSeekerModel = validator.New()
var validateHireModel = validator.New()

func CreateJobSeeker(c *fiber.Ctx) error {
	fmt.Print("Jobs")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var jobseeker models.JobSeeker
	defer cancel()
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("image upload error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	// generate new uuid for image name
	uniqueId := uuid.New()

	// remove "- from imageName"

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// extract image extension from original file filename

	fileExt := strings.Split(file.Filename, ".")[1]

	// generate image from filename and extension
	image := fmt.Sprintf("%s.%s", filename, fileExt)

	// save image to ./images dir
	err = c.SaveFile(file, fmt.Sprintf("./images/%s", image))

	if err != nil {
		log.Println("image save error --> ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	// generate image url to serve to client using CDN

	imageUrl := fmt.Sprintf("http://localhost:8080/images/%s", image)

	//validate the request body
	if err := c.BodyParser(&jobseeker); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&jobseeker); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newJobseeker := models.JobSeeker{
		Id:       primitive.NewObjectID(),
		Name:     jobseeker.Name,
		Email:    jobseeker.Email,
		Skills:   jobseeker.Skills,
		Phone:    jobseeker.Phone,
		Location: jobseeker.Location,
		Exp:      jobseeker.Exp,
		Salary:   jobseeker.Salary,
		Desc:     jobseeker.Desc,
		FileName: imageUrl,
	}
	fmt.Println("newJobseeker", newJobseeker)
	result, err := JobSeekerCollection.InsertOne(ctx, newJobseeker)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func CreateHire(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var hire models.Hire
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&hire); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&hire); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newHire := models.Hire{
		Id:          primitive.NewObjectID(),
		Name:        hire.Name,
		Email:       hire.Email,
		Skills:      hire.Skills,
		Phone:       hire.Phone,
		Location:    hire.Location,
		Exp:         hire.Exp,
		Salary:      hire.Salary,
		Jobtype:     hire.Jobtype,
		Companyname: hire.Companyname,
		Jobtitle:    hire.Jobtitle,
		Jobopening:  hire.Jobopening,
	}

	result, err := HireCollection.InsertOne(ctx, newHire)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})

	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}
