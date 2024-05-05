package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoDbConnectionString = "mongodb://localhost:27017"
const databaseName = "Netflix"

var Db *mongo.Database

func init() {

	clientOption := options.Client().ApplyURI(mongoDbConnectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb connected successfully")

	Db = client.Database(databaseName)
}
