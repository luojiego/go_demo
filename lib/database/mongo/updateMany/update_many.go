package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"time"
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

	PlayerList []userItem `json:"player_list" bson:"playerList"`
}

type userItem struct {
	UserId   int   `json:"user_id" bson:"user_id"`
	JoinTime int64 `json:"join_time" bson:"join_time"`
}

func getUpdate(id int) []userItem {
	var arr []userItem
	for i := 0; i < rand.Intn(30)+1; i++ {
		arr = append(arr, userItem{
			UserId:   id*100 + i,
			JoinTime: time.Now().Unix(),
		})
	}
	return arr
}

func main() {
	now := time.Now()
	/* 批量更新 _id 10000 - 11000 的数据
	for i := 10000; i < 11000; i++ {
		filter := bson.D{{"_id", i}}
		update := bson.M{"$set": bson.D{{"playerList",getUpdate(i)}}}
		collection.UpdateMany(ctx, filter, update)
	}*/

	/*
		查询 1099700 是否在用户列表中
		for i := 0; i < 1000; i++ {
			collection.FindOne(ctx, bson.D{{"playerList.user_id", 1099700}})
		}

		//查询结果为 5.602s
	*/

	/*
		增加索引，再次查询
		db.users1.createIndex({playerList:1})
		for i := 0; i < 1000; i++ {
			collection.FindOne(ctx, bson.D{{"playerList.user_id", 1099700}})
		}

		//没有作用
	*/

	/*
		重新增加索引，再次查询
		db.users1.createIndex({"playerList.userId":1})
	*/
	/*for i := 0; i < 1000; i++ {
		collection.FindOne(ctx, bson.D{{"playerList.user_id", 1099700}})
	}

	*/
	//0.469s 正确的增加索引，可以提交查询速度

	/*arr := []int{
		1000,1005,10023, 1999, 10025, 10006, 10009, 10333, 10888, 10999,
		1001,1006,10025, 1998, 10022, 10009, 10019, 10533, 10898, 10219,
		1002,1007,10024, 1997, 10026, 10008, 10039, 10833, 10668, 10839,
	}

	sort.Ints(arr)

	fmt.Println("arr: ", arr, "len: ", len(arr))
	cursor, err := collection.Find(ctx, bson.M{"_id": bson.M{"$in": arr}})
	if err != nil {
		panic(err)
	}
	for cursor.Next(ctx) {
		d := Data{}
		cursor.Decode(&d)
		fmt.Println(d)
	}
	defer cursor.Close(ctx)*/

	//未查找到 则插入一条新的数据
	d := Data{
		Id:       20001,
		Name:     "我最叼",
		Level:    2,
		VipLevel: 3,
		Trophy:   2000,
		//PlayerList: nil,
	}
	op := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	r := collection.FindOneAndUpdate(ctx, bson.M{"name": "user20001"}, bson.M{"$set": d}, op)

	d1 := &Data{}
	err := r.Decode(d1)
	if err != nil {
		panic(err)
	}

	fmt.Println(d1)

	fmt.Printf("%0.5f\n", time.Since(now).Seconds())
	//collection.InsertMany(ctx, arr)
}
