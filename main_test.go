package main

import (
	"context"
	"testing"
	"fmt"
	"github.com/wi7sonjoseph/mongodb_go_test/service"
)

func TestMongoDB(t *testing.T) {
	client := service.SetupDB()

	db := client.Database("test")
	collection := db.Collection("trainers")
	
	//Document to Insert
	brock := service.Trainer{"Wilson", 15, "Pewter City"}
	insertResult, err := collection.InsertOne(context.TODO(), brock)
	if err != nil {
		t.Errorf("Failed to Insert:")
		t.Errorf(err.Error())
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)
	
	/*err = collection.Drop(context.TODO())
	if err != nil {
		t.Errorf("Failed to Drop Collection:")
		t.Errorf(err.Error())
	}
	
	err = db.Drop(context.TODO())
	if err != nil {
		t.Errorf("Failed to Drop Database:")
		t.Errorf(err.Error())
	}*/
}
