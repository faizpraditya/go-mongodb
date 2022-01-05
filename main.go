package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb://localhost:27017"

func main() {
	credential := options.Credential{
		Username: "jack",
		Password: "12345678",
	}

	clientOptions := options.Client()
	// Udah buat, maka dari itu di SetAuth(credential)
	clientOptions.ApplyURI(uri).SetAuth(credential)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully to connect and pinged")

	// coll := client.Database("db_enigma").Collection("students")
	coll2 := client.Database("db_enigma").Collection("products")

	// Insert for Students
	// const layout = "2006-01-02"
	// dt, _ := time.Parse(layout, "2022-01-05")
	// newStudent := Student{
	// 	Id:       primitive.NewObjectID(),
	// 	Name:     "Phil Foden",
	// 	Gender:   "M",
	// 	Age:      23,
	// 	JoinDate: dt,
	// 	IdCard:   "111",
	// 	Senior:   true,
	// }

	// InsertOneStudent(ctx, coll, newStudent)

	// FindAllStudent(ctx, coll)

	// FindStudentByGenderAndAge(ctx, coll, "M", 23)

	// CountStudent(ctx, coll)
	// CountStudentByAge(ctx, coll, 23)

	// FindStudentStructByGenderAndAge(ctx, coll, "M", 23)

	CountProductByCategory(ctx, coll2, "food")
}
