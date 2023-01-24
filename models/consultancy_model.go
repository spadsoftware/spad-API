package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type JobSeeker struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Skills   string             `json:"skills,omitempty" validate:"required"`
	Phone    string             `json:"phone,omitempty" validate:"required"`
	Location string             `json:"location,omitempty" validate:"required"`
	Exp      string             `json:"exp,omitempty" validate:"required"`
	Salary   string             `json:"salary,omitempty" validate:"required"`
	Desc     string             `json:"desc,omitempty" validate:"required"`
	FileName string             `json:"fileName,omitempty" validate:"required"`
}

type Hire struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Name        string             `json:"name,omitempty" validate:"required"`
	Email       string             `json:"email,omitempty" validate:"required"`
	Skills      string             `json:"skills,omitempty" validate:"required"`
	Phone       string             `json:"phone,omitempty" validate:"required"`
	Location    string             `json:"location,omitempty" validate:"required"`
	Exp         string             `json:"exp,omitempty" validate:"required"`
	Salary      string             `json:"salary,omitempty" validate:"required"`
	Jobtype     string             `json:"jobtype,omitempty" validate:"required"`
	Companyname string             `json:"companyname,omitempty" validate:"required"`
	Jobtitle    string             `json:"jobtitle,omitempty" validate:"required"`
	Jobopening  string             `json:"jobopening,omitempty" validate:"required"`
}
