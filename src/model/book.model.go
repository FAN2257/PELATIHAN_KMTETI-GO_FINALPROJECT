package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title  string            `bson:"title"`
	Author string            `bson:"author"`
	Price int                `bson:"price"`
	Stock int                `bson:"stock"`
}