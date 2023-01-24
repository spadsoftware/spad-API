package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Subject  string             `json:"subject,omitempty" validate:"required"`
	Message  string             `json:"message,omitempty" validate:"required"`
	EmailId    string             `json:"email,omitempty" validate:"required"`
	PhoneNum    string             `json:"phone,omitempty" validate:"required"`
	Category string             `json:"category,omitempty" validate:"required"`
}

type Appoinment struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	DateTime  string             `json:"dateTime,omitempty" validate:"required"`
	Message  string             `json:"message,omitempty" validate:"required"`
	EmailId    string             `json:"email,omitempty" validate:"required"`
	PhoneNum    string             `json:"phone,omitempty" validate:"required"`
	Category string             `json:"category,omitempty" validate:"required"`
}