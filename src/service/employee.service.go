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

func GetAllEmployee() (*model.EmployeeResponse, error) {
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

	var prodList []*model.EmployeeGetAll

	for cur.Next(context.TODO()) {
		var prod model.Employee
		cur.Decode(&prod)
		prodList = append(prodList, &model.EmployeeGetAll{
			Name:     prod.Name,
			JoinDate: prod.JoinDate,
			Status:   prod.Status,
		})
	}
	return &model.EmployeeResponse{
		Data: prodList,
	}, nil
}

func CreateEmployee(req io.Reader) error {
	var empReq model.EmployeeRequest
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

	if empReq.Status != model.EmployeeKontrak && empReq.Status != model.EmployeeTetap {
		return errors.New("bad request")
	}

	coll := db.MongoDB.Collection("employee")
	_, err = coll.InsertOne(context.TODO(), model.Employee{
		ID:            primitive.NewObjectID(),
		Name:          empReq.Name,
		NIK:           empReq.NIK,
		LastEducation: empReq.LastEducation,
		JoinDate:      empReq.JoinDate,
		Status:        empReq.Status,
	})
	if err != nil {
		log.Default().Println(err.Error())
		return errors.New("internal server error")
	}
	return nil
}
