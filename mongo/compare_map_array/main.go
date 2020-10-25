package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"strconv"
	"time"
)

type Data struct {
	Id      int
	Name    string
	Age     int8
	Address string
}

type Data1 struct {
	Name    string
	Age     int8
	Address string
}

type Array struct {
	Id   int `json:"id" bson:"_id"`
	Data []Data
}

type Map struct {
	Id   int `json:"id" bson:"_id"`
	Data map[int]Data1
}

var (
	ctx             = context.Background()
	collectionArray *mongo.Collection
	collectionMap   *mongo.Collection
)

func init() {
	ctx = context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.196.16:27017"))
	if err != nil {
		panic(err)
	}

	collectionArray = client.Database("test").Collection("test_array")
	collectionMap = client.Database("test").Collection("test_map")
}

func InsertArray(arr *Array) {
	collectionArray.InsertOne(ctx, arr)
}

func InsertMap(m *Map) {
	collectionMap.InsertOne(ctx, m)
}

func updateArray(id int, age int8) float64 {
	t := time.Now()
	if _, err := collectionArray.UpdateOne(ctx, bson.M{"_id": 9876}, bson.D{{"$set",
		bson.M{"data.$[item].age": age}}}, options.Update().SetArrayFilters(options.ArrayFilters{Filters: bson.A{bson.M{"item.id": id}}})); err != nil {
		fmt.Println(err)
	}
	return time.Since(t).Seconds()
}

func updateMap(id int, age int8) float64 {
	t := time.Now()
	if _, err := collectionMap.UpdateOne(ctx, bson.M{"_id": 9876}, bson.M{"$set": bson.M{"data." + strconv.Itoa(id) + ".age": age}}); err != nil {
		fmt.Println(err)
	}
	return time.Since(t).Seconds()
}

func insert() {
	arr := &Array{Id: 9876}
	m := &Map{
		Id:   9876,
		Data: make(map[int]Data1),
	}
	for i := 0; i < 1000; i++ {
		arr.Data = append(arr.Data, Data{
			Id:      i + 1,
			Name:    "test" + strconv.Itoa(i+1),
			Age:     int8(rand.Intn(40)) + 1,
			Address: "天堂路" + strconv.Itoa(i+1) + "号",
		})

		m.Data[i+1] = Data1{
			Name:    "test" + strconv.Itoa(i+1),
			Age:     int8(rand.Intn(40)) + 1,
			Address: "天堂路" + strconv.Itoa(i+1) + "号",
		}
	}
	InsertArray(arr)
	InsertMap(m)
}

func main() {
	var a, b float64
	for i := 0; i < 1000; i++ {
		id := rand.Intn(1000) + 1
		age := int8(rand.Intn(50) + 1)
		a += updateArray(id, age)
		b += updateMap(id, age)
	}
	fmt.Printf("a: %.2f\n", a)
	fmt.Printf("b: %.2f\n", b)
}
