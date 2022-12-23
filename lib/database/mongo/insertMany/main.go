package main

import (
	"context"
	"math/rand"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//使用多个 goroutine 来 inc 一个value
var (
	ctx        = context.Background()
	collection *mongo.Collection
)

func init() {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.196.20:27017"))
	if err != nil {
		panic(err)
	}

	collection = client.Database("test").Collection("users1")
}

type Data struct {
	Id       int    `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Level    int8   `json:"level" bson:"level"`
	VipLevel int8   `json:"vip_level" bson:"vip_level"`
	Trophy   int    `json:"trophy" bson:"trophy"`
}

func main() {
	var arr []interface{}
	for i := 1000; i < 11000; i++ {
		arr = append(arr, Data{
			Id:       i,
			Name:     "user" + strconv.FormatInt(int64(i), 10),
			Level:    int8(rand.Intn(128)),
			VipLevel: int8(rand.Intn(128)),
			Trophy:   rand.Intn(65535),
		})
	}

	// collection.InsertMany(ctx, arr)
	// collection.InsertOne(context.Background(), bson.M{"_id": 1, "val": []int{1, 2, 3}})

	collection.UpdateOne(context.Background(), bson.M{"_id": 1}, bson.M{"$push": bson.M{"val": bson.M{"$each": []int{4, 5}}}})
}
