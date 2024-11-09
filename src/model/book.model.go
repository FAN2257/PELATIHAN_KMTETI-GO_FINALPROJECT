package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Binary JSON (BSON)
type Book struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Author      string             `bson:"author"`
	DateRelease string             `bson:"date_release"`
	Price       int                `bson:"price"`
	Stock       int                `bson:"stock"`
}

// Post Struct
type BookRequest struct {
	ID          primitive.ObjectID `json:"_id,omitempty"`
	Title       string             `json:"title"`
	Author      string             `json:"author"`
	DateRelease string             `json:"date_release"`
	Price       int                `json:"price"`
	Stock       int                `json:"stock"`
}

// Get Struct All Book
type BookGetAll struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

// Response Struct
type BookResponse struct {
	Data []*BookGetAll `json:"data"`
}

// Get Struct Book Detail
type BookID struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	DateRelease string `json:"date_release"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}
