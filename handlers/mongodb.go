package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mdb *mongo.Database
var ctx context.Context
var ctxcancel context.CancelFunc

var mongoURI string = "mongodb://ip:port"
var mongoDBName string = "documentname"

func mongoConn(dbname, uri string) (mdb *mongo.Database) {
	ctx, ctxcancel = context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err.Error())
	}
	mdb = client.Database(dbname)
	return
}

func mdbFind(collection string, filter bson.D) (result []interface{}) {
	cursor, err := mdb.Collection(collection).Find(ctx, filter)
	if err != nil {
		panic(err.Error())
	}
	for cursor.Next(ctx) {
		var row bson.M
		err := cursor.Decode(&row)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, row)
	}
	return
}

func mdbUpdateOne(collection string, filter, update bson.D) (err error) {
	_, err = mdb.Collection(collection).UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	return
}

func mdbInsert(collection string, insert []interface{}) (insertedID []interface{}, err error) {
	insertManyResult, err := mdb.Collection(collection).InsertMany(ctx, insert)
	insertedID = insertManyResult.InsertedIDs
	return
}

func main() {
	mdb = mongoConn(mongoDBName, mongoURI)
	defer ctxcancel()

	// Find filter
	filter := bson.D{{Key: "section", Value: "dashboard"}}

	// Find example
	result := mdbFind("collectionname", filter)
	for idx, row := range result {
		fmt.Printf("Row %v ID %v:\nRow: %v\n", idx+1, row.(bson.M)["_id"], row)
	}

	// UpdateOne if exist else insert example
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "updated", Value: true}, {Key: "timestamp", Value: time.Now().Unix()}}}}
	err := mdbUpdateOne("collectionname", filter, update)
	if err != nil {
		panic(err.Error())
	}

	// InsertMany example

	insert := []interface{}{bson.D{{Key: "test_key1_1", Value: "1_1"}, {Key: "test_key1_2", Value: "1_2"}}, bson.D{{Key: "test_key2", Value: "2"}}}
	IDs, err := mdbInsert("collectionname", insert)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Inserted IDs: %v", IDs)
}
