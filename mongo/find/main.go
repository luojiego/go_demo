package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//docker run -p 27017:27017 -v /home/dev:/data/db --name mongo --restart always -d mongo

func Find(ctx context.Context, collection *mongo.Collection) {
	fmt.Println("---------------------------Find------------------------------")
	cur, err := collection.Find(ctx, bson.D{}, options.Find().SetProjection(bson.M{"age": 1}))
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

type Info struct {
	Address     string `json:"address" bson:"address"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
}

type Data struct {
	Name string `json:"name" bson:"name"`
	Info Info
}

func insert(ctx context.Context, collection *mongo.Collection) error {
	d := Data{
		Name: "罗杰",
		Info: Info{
			Address:     "陕西省西安市",
			PhoneNumber: "1539900000",
		},
	}

	one, err := collection.InsertOne(ctx, d)
	if err != nil {
		return err
	}
	fmt.Println(one.InsertedID)
	return nil
}

func find1(ctx context.Context, collection *mongo.Collection) (info Data, err error) {
	objectId, _ := primitive.ObjectIDFromHex("5ecf97fbd9c5c823e699eeb4")
	err = collection.FindOne(ctx,
		bson.D{{"_id", objectId}},
		options.FindOne().SetProjection(bson.D{{"info", 1}, {"_id", 0}}),
	).Decode(&info)
	fmt.Println(info)
	return info, err
}

var (
	ctx        = context.Background()
	collection *mongo.Collection
)

func init() {
	ctx = context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.196.50:27017"))
	if err != nil {
		panic(err)
	}

	//fmt.Println(client)

	collection = client.Database("test").Collection("user3")
}

func main() {

	// Find(ctx, collection)

	//FindOne(ctx, collection)

	//insert(ctx, collection)

	//find1(ctx, collection)

	//insertMany(ctx, collection)

	//findTest1()

	// insertOne()

	findId(321,1234)
}

type Item struct {
	UserId int `json:"user_id" bson:"user_id"`
	List []int `json:"list" bson:"list"`
}

type Data1 struct {
	Id int `json:"id" bson:"_id"`
	List []Item `json:"list" bson:"list"`
}

func findId(id1, id2 int) {
	fmt.Println(id1, id2)

	count, err1 := collection.CountDocuments(ctx, bson.M{"_id": id1, "list.user_id": id2})
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println("count: ", count)

	collection.

	one := collection.FindOne(ctx, bson.D{{"_id", id1}, {"list.user_id", id2}})
	if one.Err() != nil {
		fmt.Println("没有找到")
	} else {
		fmt.Println("我找到了")
	}
	/*
	find, err := collection.Find(ctx,
		bson.M{"_id": id1, "list.user_id": id2})
		// options.Find().SetProjection(bson.M{"_id": 1}))
	if err != nil {
		fmt.Println(err)
		return
	}

	// data := &Data1{}
	var data bson.M
	if err := find.Decode(&data); err != nil {
		fmt.Printf("find decode err: %s\n", err)
		return
	}
	fmt.Println(data)

 */

}

func findTest() {
	cur, err := collection.Find(
		ctx,
		bson.D{
			/*{"instock", bson.D{
				{"qty", 5},
				{"warehouse", "A"},
			}},*/
			{"instock.warehouse", "A"},
		})

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal("Decode: ", err)
		}
		// do something with result....
		fmt.Println("result: ", result)
	}
	//r := bson.D{}
	//err = cursor.Decode(&r)

	fmt.Println(err)
}

func findTest1() {
	one := collection.FindOne(
		ctx,
		bson.M{
			/*{"instock", bson.D{
				{"qty", 5},
				{"warehouse", "A"},
			}},*/
			"instock.warehouse": "A",
		})

	r := bson.D{}
	err := one.Decode(&r)

	fmt.Println(err)
}

func insertOne() {
	docs := &bson.D{{"item", "journal"}, {"instock", bson.A{
		bson.D{
			{"warehouse", "A"},
			{"qty", 5},
		},
		bson.D{
			{"warehouse", "C"},
			{"qty", 15},
		},
	}}}
	one, err := collection.InsertOne(context.Background(), docs)
	if err != nil {
		panic(err)
	}

	s := one.InsertedID.(primitive.ObjectID).Hex()
	fmt.Println(s)
}

func insertMany(ctx context.Context, collection *mongo.Collection) {
	docs := []interface{}{
		bson.D{
			{"item", "journal"},
			{"instock", bson.A{
				bson.D{
					{"warehouse", "A"},
					{"qty", 5},
				},
				bson.D{
					{"warehouse", "C"},
					{"qty", 15},
				},
			}},
		},
		bson.D{
			{"item", "notebook"},
			{"instock", bson.A{
				bson.D{
					{"warehouse", "C"},
					{"qty", 5},
				},
			}},
		},
		bson.D{
			{"item", "paper"},
			{"instock", bson.A{
				bson.D{
					{"warehouse", "A"},
					{"qty", 60},
				},
				bson.D{
					{"warehouse", "B"},
					{"qty", 15},
				},
			}},
		},
		bson.D{
			{"item", "planner"},
			{"instock", bson.A{
				bson.D{
					{"warehouse", "A"},
					{"qty", 40},
				},
				bson.D{
					{"warehouse", "B"},
					{"qty", 5},
				},
			}},
		},
		bson.D{
			{"item", "postcard"},
			{"instock", bson.A{
				bson.D{
					{"warehouse", "B"},
					{"qty", 15},
				},
				bson.D{
					{"warehouse", "C"},
					{"qty", 35},
				},
			}},
		},
	}

	collection.InsertMany(context.Background(), docs)
}
