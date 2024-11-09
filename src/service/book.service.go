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

type BookID struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
	Stock  int    `json:"stock"`
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

	coll := db.MongoDB.Collection("book")
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

func GetBookByID(id string) (*BookID, error) {
	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("book")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("bad request")
	}

	var book model.Book
	err = coll.FindOne(context.TODO(), bson.D{{Key: "_id", Value: objID}}).Decode(&book)
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}

	return &BookID{
		Title:  book.Title,
		Author: book.Author,
		Price:  book.Price,
		Stock:  book.Stock,
	}, nil
}

func UpdateBook(id string, req io.Reader) error {
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
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("bad request")
	}

	_, err = coll.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: objID}}, bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: bookReq.Title},
			{Key: "author", Value: bookReq.Author},
			{Key: "price", Value: bookReq.Price},
		}},
	})
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}
	return nil
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

func DeleteBookByID(id string) error {
	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("book")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("bad request")
	}

	_, err = coll.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: objID}})
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}
	return nil
}
