package golang_mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-mongodb/helper"
	"golang-mongodb/model"
	"testing"
)

func TestCountDocument(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	fmt.Println(coll.CountDocuments(context.Background(), bson.D{}))
}

func TestLimit(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	opt := options.Find()
	opt.SetLimit(4)

	var model []model.Product
	cur, err := coll.Find(context.Background(), bson.D{}, opt)
	cur.All(context.Background(), &model)
	helper.PanicIfError(err)
	fmt.Println(model)
}
func TestSkip(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	opt := options.Find()
	opt.SetSkip(2)

	var model []model.Product
	cur, err := coll.Find(context.Background(), bson.D{}, opt)
	cur.All(context.Background(), &model)
	helper.PanicIfError(err)
	fmt.Println(model)
}
func TestSkipLimit(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	opt := options.Find()
	opt.SetSkip(2)
	opt.SetLimit(4)

	var model []model.Product
	cur, err := coll.Find(context.Background(), bson.D{}, opt)
	cur.All(context.Background(), &model)
	helper.PanicIfError(err)
	fmt.Println(model)
}
func TestSort(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	opt := options.Find()
	opt.SetSort(bson.D{{"name", 1}, {"category", -1}})

	var model []model.Product
	cur, err := coll.Find(context.Background(), bson.D{}, opt)
	cur.All(context.Background(), &model)
	helper.PanicIfError(err)
	fmt.Println(model)
}
