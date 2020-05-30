package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"strconv"
)

//使用多个 goroutine 来 inc 一个value
var (
	ctx        = context.Background()
	collection *mongo.Collection
)

func init() {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.196.50:27017"))
	if err != nil {
		panic(err)
	}

	collection = client.Database("game").Collection("users1")
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

	collection.InsertMany(ctx, arr)
}
