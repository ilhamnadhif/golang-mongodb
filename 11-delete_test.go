package golang_mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"golang-mongodb/helper"
	"golang-mongodb/model"
	"testing"
)

func TestInsertCustomer(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("customers")

	res, err := coll.InsertOne(context.Background(), model.Customer{
		Id:       "spammer",
		FullName: "Spammer",
	})
	helper.PanicIfError(err)
	fmt.Println(res)
}
func TestDeleteOne(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("customers")

	res, err := coll.DeleteOne(context.Background(), bson.D{{"_id", "spammer"}})
	helper.PanicIfError(err)
	fmt.Println(res)
}
func TestInsertManyCustomer(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("customers")

	docs := []interface{}{
		model.Customer{
			Id:       "spammer1",
			FullName: "Spammer1",
		},
		model.Customer{
			Id:       "spammer2",
			FullName: "Spammer2",
		},
		model.Customer{
			Id:       "spammer3",
			FullName: "Spammer3",
		},
	}
	res, err := coll.InsertMany(context.Background(), docs)
	helper.PanicIfError(err)
	fmt.Println(res)
}
func TestDeleteMany(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("customers")

	res, err := coll.DeleteMany(context.Background(), bson.D{
		{"_id", bson.D{{
			"$regex", "spammer",
		}}},
	})
	helper.PanicIfError(err)
	fmt.Println(res)
}
