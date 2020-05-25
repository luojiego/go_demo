package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//docker run -p 27017:27017 -v /home/dev:/data/db --name mongo --restart always -d mongo

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.196.50:27017"))
	if err != nil {
		panic(err)
	}

	fmt.Println(client)
	ctx := context.Background()
	collection := client.Database("game").Collection("user")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result....
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
