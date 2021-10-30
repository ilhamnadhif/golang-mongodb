package golang_mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang-mongodb/helper"
	"golang-mongodb/model"
	"testing"
)

func TestFindOne(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("customers")
	var customer model.Customer
	err := coll.FindOne(context.Background(), bson.D{{"_id", "khannedy"}}).Decode(&customer)
	helper.PanicIfError(err)
	fmt.Println("Id", customer.Id)
	fmt.Println("Name", customer.Name)

}
func TestFindOneNestedArray(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("orders")
	var order model.Order
	err := coll.FindOne(context.Background(), bson.D{{"items.product_id", 1}}).Decode(&order)
	helper.PanicIfError(err)
	fmt.Println(order)
}
func TestFindOneByObjectId(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("orders")

	obejectId, errorr := primitive.ObjectIDFromHex("617c99a5f7e1b689f781e67c")
	helper.PanicIfError(errorr)

	var order model.Order
	err := coll.FindOne(context.Background(), bson.D{{"_id", obejectId}}).Decode(&order)
	helper.PanicIfError(err)
	fmt.Println(order)
}

func TestFind(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	cur, err := coll.Find(context.Background(), bson.D{})
	helper.PanicIfError(err)

	var products []model.Product
	err2 := cur.All(context.Background(), &products)
	helper.PanicIfError(err2)

	fmt.Println(products)
}
