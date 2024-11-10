package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeStatus string

const (
	EmployeeKontrak EmployeeStatus = "KONTRAK"
	EmployeeTetap   EmployeeStatus = "TETAP"
)

// Binary JSON (BSON)
type Employee struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name"`
	NIK           int                `bson:"nik"`
	LastEducation string             `bson:"last_education"`
	JoinDate      time.Time          `bson:"join_date"`
	Status        EmployeeStatus     `bson:"status"`
}

// Post Struct an Employee
type EmployeeRequest struct {
	ID            primitive.ObjectID `json:"_id,omitempty"`
	Name          string             `json:"name"`
	NIK           int                `json:"nik"`
	LastEducation string             `json:"last_education"`
	JoinDate      time.Time          `json:"join_date"`
	Status        EmployeeStatus     `json:"status"`
}

// Get Struct All Employee
type EmployeeGetAll struct {
	Name     string         `json:"name"`
	JoinDate time.Time      `json:"join_date"`
	Status   EmployeeStatus `json:"status"`
}

type EmployeeResponse struct {
	Data []*EmployeeGetAll `json:"data"`
}
