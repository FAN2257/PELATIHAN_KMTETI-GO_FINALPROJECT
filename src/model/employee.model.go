package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name"`
	NIK           int                `bson:"nik"`
	LastEducation string             `bson:"last_education"`
	JoinDate      int32              `bson:"join_date"`
	Status        bool               `bson:"status"`
}
