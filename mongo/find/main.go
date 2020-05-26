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

func Find(ctx context.Context, collection *mongo.Collection) {
	fmt.Println("---------------------------Find------------------------------")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal("Find: ", err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal("Decode: ", err)
		}
		// do something with result....
		fmt.Println("result: ", result)
		fmt.Println("name: ", cur.Current.Lookup("name"))

		d := &struct {
			Id   string `json:"_id" bson:"_id"`
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{}

		err = cur.Decode(d)
		if err != nil {
			log.Println("Decode2 :", err)
		}
		fmt.Println(d)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("---------------------------Find------------------------------")
}

func FindOne(ctx context.Context, collection *mongo.Collection) {
	fmt.Println("---------------------------FindOne------------------------------")
	one := collection.FindOne(ctx, collection)
	bytes, err := one.DecodeBytes()
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes.String())
	fmt.Println("name: ", bytes.Lookup("name"), " _id:", bytes.Lookup("_id"))
	fmt.Println("---------------------------FindOne------------------------------")
}

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.196.50:27017"))
	if err != nil {
		panic(err)
	}

	fmt.Println(client)

	collection := client.Database("game").Collection("users")
	Find(ctx, collection)

	FindOne(ctx, collection)
}
