package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pages struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Pid      string             `json:"pid,omitempty" validate:"required"`
	Ptitle   string             `json:"ptitle,omitempty" validate:"required"`
	Pdesc    string             `json:"pdesc,omitempty" validate:"required"`
	Pkeyword string             `json:"pkeyword,omitempty" validate:"required"`
	Pimg     string             `json:"pimg,omitempty" validate:"required"`
	PimgAlt  string             `json:"pimgAlt,omitempty" validate:"required"`
	Pauthor  string             `json:"pauthor,omitempty" validate:"required"`
}

// type Page []Pages
