package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name"`
	NIK           int                `bson:"nik"`
	LastEducation string             `bson:"last_education"`
	JoinDate      time.Time          `bson:"join_date"`
	Status        bool               `bson:"status"`
}