package data

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
)

// Client declare the database client
var Client *mongo.Client

// Ctx declare the database context
var Ctx = context.TODO()

// DBConnect Database Connect module
func DBConnect(l *log.Logger) *mongo.Client {
	l.Println("********* Trying to connect with MongoDB")
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(Ctx, clientOption)
	Client = client
	if err != nil {
		l.Println("Error connecting the database: %v", err)
		return nil
	}

	err = Client.Ping(Ctx, nil)
	if err != nil {
		l.Println("Database Connection is unsuccessful")
		return nil
	}
	l.Println("********* Connection successful with MongoDB")

	return Client
}

// GetDB CreateDB is used to create a database with desired table name
func GetDB(l *log.Logger) (*mongo.Collection, context.Context) {
	l.Println("********* Trying to create table with MongoDB")
	l.Printf("Client: %v", Client)
	l.Printf("Ctx: %v", Ctx)
	return Client.Database("testDB").Collection("persons"), Ctx
}
