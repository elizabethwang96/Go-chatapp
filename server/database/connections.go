package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var userCollection *mongo.Collection

func Connect() {

	connection, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://elizabethwang:20000906@cluster0.mf7bkze.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	if err := connection.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Print("Connect succeeded")
	userCollection := connection.Database("Go_Chatapp").Collection("users")
}
