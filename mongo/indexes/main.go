package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"time"
)

//使用多个 goroutine 来 inc 一个value
var (
	ctx        = context.Background()
	collection *mongo.Collection
	//将 name 为 user1980 的用户 trophy 通过 1000 个 goroutine 各增加1
	//find = bson.D{{"name", "user1980"}}
	//update = bson.D{{"$inc", bson.M{"trophy":1}}}

	//将 name 为 user10025 的用户 trophy 通过 1000 个 goroutine 各增加1
	//find = bson.D{{"name", "user10025"}}
	//update = bson.D{{"$inc", bson.M{"trophy":1}}}

	//将 _id 为 10025 的用户 trophy 通过 1000 个 goroutine 各增加1
	//objectId, _ = primitive.ObjectIDFromHex("10025")
	find   = bson.D{{"_id", 10025}}
	update = bson.D{{"$inc", bson.M{"trophy": 1}}}
)

func init() {
	//ctx = context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.196.50:27017"))
	if err != nil {
		panic(err)
	}

	collection = client.Database("game").Collection("users1")
}

func query(wg *sync.WaitGroup) {
	defer wg.Done()
	collection.FindOne(ctx, find, options.FindOne().SetProjection(bson.M{"trophy": 1}))
}

func main() {
	cursor, err := collection.Indexes().List(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Get a slice of all indexes returned and print them out.
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)

	now := time.Now()
	wg := sync.WaitGroup{}
	count := 1000
	for i := 0; i < count; i++ {
		wg.Add(1)
		go query(&wg)
	}
	wg.Wait()
	fmt.Printf("update count: %d use: %0.5fs\n", count, time.Since(now).Seconds())

}
