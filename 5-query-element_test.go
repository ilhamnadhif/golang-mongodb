package golang_mongodb

import (
	"golang-mongodb/helper"
	"golang-mongodb/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestFindExists(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	cur, err := coll.Find(context.Background(), bson.D{
		{
			"category", bson.D{
				{
					"$exists", false,
				},
			},
		},
	})
	helper.PanicIfError(err)

	var products []model.Product
	err2 := cur.All(context.Background(), &products)
	helper.PanicIfError(err2)

	fmt.Println(products)
}

func TestFindType(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	cur, err := coll.Find(context.Background(), bson.D{
		{
			"category", bson.D{
				{
					"$type", "string",
				},
			},
		},
	})
	helper.PanicIfError(err)

	var products []model.Product
	err2 := cur.All(context.Background(), &products)
	helper.PanicIfError(err2)

	fmt.Println(products)
}

func TestFindTypeIn(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	cur, err := coll.Find(context.Background(), bson.D{
		{
			"price", bson.D{
			{
				"$type", []string{"int", "long"},
			},
		},
		},
	})
	helper.PanicIfError(err)

	var products []model.Product
	err2 := cur.All(context.Background(), &products)
	helper.PanicIfError(err2)

	fmt.Println(products)
}
