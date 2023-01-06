package main

import (
	"context"
	"fmt"
	"insert/testproto"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	structcodec, _ := bsoncodec.NewStructCodec(myStructTagParser)

	rb := bson.NewRegistryBuilder()
	rb.RegisterDefaultEncoder(reflect.Struct, structcodec)
	rb.RegisterDefaultDecoder(reflect.Struct, structcodec)

	client, err := mongo.Connect(ctx, options.Client().SetRegistry(rb.Build()).ApplyURI("mongodb://192.168.196.17:27018"))
	if err != nil {
		panic(err)
	}

	/*
		//fmt.Println(client)
		collection := client.Database("game").Collection("friends")
		//insertTest(ctx, collection)

		insertFriends(ctx, collection)
	*/

	t := testproto.Person{
		Name:        "罗杰",
		AddressList: []string{"中国", "陕西", "西安", "新城区"},
	}

	collection := client.Database("test").Collection("person")

	result, err := collection.InsertOne(context.Background(), &t)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	res := &testproto.Person{}
	id, _ := primitive.ObjectIDFromHex("63b7fbc904fb09d43a9564f9")
	if err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(res); err != nil {
		panic(err)
	}
	fmt.Println(res)

	u := &testproto.User{
		Id:   1000,
		Name: "罗杰",
	}

	userCollection := client.Database("test").Collection("user")
	if _, err := userCollection.InsertOne(context.Background(), u); err != nil {
		panic(err)
	}
}

var myStructTagParser bsoncodec.StructTagParserFunc = func(sf reflect.StructField) (bsoncodec.StructTags, error) {
	key := strings.ToLower(sf.Name[:1]) + sf.Name[1:]
	tag, ok := sf.Tag.Lookup("bson")
	if !ok && !strings.Contains(string(sf.Tag), ":") && len(sf.Tag) > 0 {
		tag = string(sf.Tag)
	}
	return parseTags(key, tag)
}

func parseTags(key string, tag string) (bsoncodec.StructTags, error) {
	var st bsoncodec.StructTags
	if tag == "-" {
		st.Skip = true
		return st, nil
	}

	for idx, str := range strings.Split(tag, ",") {
		if idx == 0 && str != "" {
			key = str
		}
		switch str {
		case "omitempty":
			st.OmitEmpty = true
		case "minsize":
			st.MinSize = true
		case "truncate":
			st.Truncate = true
		case "inline":
			st.Inline = true
		}
	}

	st.Name = key

	return st, nil
}
