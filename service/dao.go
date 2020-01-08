package service

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

func CrudFunctionality(client *mongo.Client) {
	db := client.Database("test")
	collection := db.Collection("trainers")
	brock := Trainer{"Mo Salah", 15, "Pewter City"}
	insertResult, err := collection.InsertOne(context.TODO(), brock)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
