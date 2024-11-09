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

type Employee struct {
	Name     string `json:"name"`
	JoinDate int32  `json:"join_date"`
	Status   bool   `json:"status"`
}

type EmployeeResponse struct {
	Data []*Employee `json:"data"`
}

type EmployeeRequest struct {
	Name     string `json:"name"`
	JoinDate int32  `json:"join_date"`
	Status   bool   `json:"status"`
}

func GetAllEmployee() (*EmployeeResponse, error) {
	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("employee")
	cur, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}

	var prodList []*Employee

	for cur.Next(context.TODO()) {
		var prod model.Employee
		cur.Decode(&prod)
		prodList = append(prodList, &Employee{
			Name:     prod.Name,
			JoinDate: prod.JoinDate,
			Status:   prod.Status,
		})
	}
	return &EmployeeResponse{
		Data: prodList,
	}, nil
}

func CreateEmployee(req io.Reader) error {
	var empReq EmployeeRequest
	err := json.NewDecoder(req).Decode(&empReq)
	if err != nil {
		return errors.New("bad request")
	}

	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	coll := db.MongoDB.Collection("employee")
	_, err = coll.InsertOne(context.TODO(), model.Employee{
		ID:            primitive.NewObjectID(),
		Name:          empReq.Name,
		NIK:           0,
		LastEducation: "",
		JoinDate:      empReq.JoinDate,
		Status:        empReq.Status,
	})
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}
	return nil
}
