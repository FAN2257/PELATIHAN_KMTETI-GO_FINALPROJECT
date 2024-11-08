package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"

	"github.com/FAN2257/PELATIHAN_KMTETI-GO_FINALPROJECT/src/db"
	"github.com/FAN2257/PELATIHAN_KMTETI-GO_FINALPROJECT/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

type BookResponse struct {
	Data []*Book `json:"data"`
}

type BookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

func GetAllBook() (*BookResponse, error) {
	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("Book")
	cur, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}

	var bookList []*Book

	for cur.Next(context.TODO()) {
		var book model.Book
		cur.Decode(&book)
		bookList = append(bookList, &Book{
			Title: book.Title,
			Price: book.Price,
		})
	}
	return &BookResponse{
		Data: bookList,
	}, nil
}

func CreateBook(req io.Reader) error {
	var bookReq BookRequest
	err := json.NewDecoder(req).Decode(&bookReq)
	if err != nil {
		return errors.New("bad request")
	}

	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("book")
	_, err = coll.InsertOne(context.TODO(), model.Book{
		ID:     primitive.NewObjectID(),
		Title:  bookReq.Title,
		Author: bookReq.Author,
		Price:  bookReq.Price,
		Stock:  0,
	})
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}
	return nil
}
