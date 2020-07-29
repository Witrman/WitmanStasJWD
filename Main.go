package main

import (
	"context"
	"fmt"
	"log"
	//  "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)

type authentication struct {
	GUID    string
	Access  string
	Refresh string
}

var client *mongo.Client

func main() {
	ConnectDB()
	addDataDB()
	closeDB()
}
func errExc(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() {
	// Create client
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(
		"mongodb+srv://user:user@clustervs.rwrgh.mongodb.net/" +
			"BaseOne?retryWrites=true&w=majority"))
	errExc(err)
	errExc(client.Connect(context.TODO()))
	errExc(client.Ping(context.TODO(), nil))
	fmt.Println("Connected to MongoDB Successful")
}

func addDataDB() {
	user1 := authentication{"Jdfjvnsskdjvns1", "GhkjhsKd89DhjivsusHhfuidh9fvhu1", "lilkdjfnlHLIUHerfgiwebllsdbLJDBLK9fvhu1"}
	user2 := authentication{"Jdfjvnsskdjvns2", "GhkjhsKd89DhjivsusHhfuidh9fvhu2", "lilkdjfnlHLIUHerfgiwebllsdbLJDBLK9fvhu2"}
	user3 := authentication{"Jdfjvnsskdjvns3", "GhkjhsKd89DhjivsusHhfuidh9fvhu3", "lilkdjfnlHLIUHerfgiwebllsdbLJDBLK9fvhu3"}

	collection := client.Database("BaseOne").Collection("ACol")
	fmt.Println("Connected to Datebase and Collection!")
	session, err := client.StartSession()
	errExc(err)

	errExc(session.StartTransaction())
	errExc(collection.Drop(context.TODO()))
	_, err = collection.InsertOne(context.TODO(), user1)
	errExc(err)
	_, err = collection.InsertOne(context.TODO(), user2)
	errExc(err)
	_, err = collection.InsertOne(context.TODO(), user3)
	errExc(err)
	errExc(session.CommitTransaction(context.TODO()))
	session.EndSession(context.TODO())

	fmt.Println("Documents dropped and added ")
}

func closeDB() {
	errExc(client.Disconnect(context.TODO()))
	fmt.Println("Connection to MongoDB closed.")
}
