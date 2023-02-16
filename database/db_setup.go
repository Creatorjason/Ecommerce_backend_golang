package database

import (
	"context" // to add deadlines to your code
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client{
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil{
		log.Fatal(err)
	}
	// Remember there is a timeout when working with database functions
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil{
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(),nil)
	if err != nil{
		log.Println("Failed to connect to database")
		return nil
	}
	fmt.Println("Successfully connected to database")

	return client
}


func ProductData(client *mongo.Client, collectionName string)*mongo.collection{

}

func UserData(client *mongo.Client, collectionName string)*mongo.Collection{

}