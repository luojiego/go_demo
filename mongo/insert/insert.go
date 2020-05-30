package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func insertMany(ctx context.Context, collection mongo.Collection) error {
	docs := []interface{}{
		bson.D{
			{"_id", 10001},
			{"name", "zz3"},
			{"age", 25},
			{"level", 2},
		},
		bson.D{
			{"_id", 10002},
			{"name", "ll4"},
			{"age", 26},
			{"level", 3},
			{"vip_level", 1},
		},
		bson.D{
			{"_id", 10003},
			{"name", "ww5"},
			{"age", 27},
			{"level", 20},
			{"vip_level", 5},
		},
		bson.D{
			{"_id", 10004},
			{"name", "yy6"},
			{"age", 28},
			{"level", 29},
			{"trophy", 22},
		},
	}
	many, err := collection.InsertMany(ctx, docs)
	if err != nil {
		return err
	}
	fmt.Println("insert: ", len(many.InsertedIDs))
	return nil
}

func insertTest(ctx context.Context, collection *mongo.Collection) error {
	d := struct {
		Id    int
		Name  string `json:"name" bson:"name"`
		Score int    `json:"score" bson:"score"`
	}{
		Id:    18973,
		Name:  "tc",
		Score: 33,
	}
	one, err := collection.InsertOne(ctx, d)
	if err != nil {
		return err
	}
	fmt.Println(one.InsertedID)
	return nil
}

func insertFriends(ctx context.Context, collection *mongo.Collection) error {
	d := struct {
		Id      int   `json:"id" bson:"_id"`
		Friends []int `json:"friends" bson:"friends"`
	}{
		Id:      10381,
		Friends: []int{10203, 10230, 13082, 14903, 15309, 10020},
	}
	one, err := collection.InsertOne(ctx, d)
	if err != nil {
		return err
	}
	fmt.Println(one.InsertedID)
	return nil
}

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.196.50:27017"))
	if err != nil {
		panic(err)
	}

	//fmt.Println(client)
	collection := client.Database("game").Collection("friends")
	//insertTest(ctx, collection)

	insertFriends(ctx, collection)
}
