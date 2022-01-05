package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// kalau ga dikasi struct tag, bakal jadi lowercase
type Student struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string
	Gender   string
	Age      int
	JoinDate time.Time `bson:"joinDate"`
	IdCard   string    `bson:"idCard"`
	Senior   bool
}
