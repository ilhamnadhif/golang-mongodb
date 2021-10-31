package golang_mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"golang-mongodb/helper"
	"golang-mongodb/model"
	"testing"
)

func TestUpdateOne(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateOne(context.Background(), bson.D{{"_id", 3}}, bson.D{
		{"$set", bson.D{{
			"category", "food",
		}}},
	})
	helper.PanicIfError(err)

	fmt.Println(res)
}
func TestUpdateMany(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateMany(context.Background(), bson.D{{
		"$and", bson.A{
			bson.D{{
				"category", bson.D{
					{
						"$eq", "food",
					},
				},
			}},
			bson.D{{
				"tags", bson.D{
					{
						"$exists", false,
					},
				},
			}},
		},
	}}, bson.D{
		{
			"$set", bson.D{
				{
					"tags", []string{"food"},
				},
			},
		},
	})
	helper.PanicIfError(err)

	fmt.Println(res)
}
func TestUpdateManyWrong(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateMany(context.Background(), bson.D{}, bson.D{
		{"$set", bson.D{{
			//"salah", "salah",
			"salah", nil,
		}}},
	})
	helper.PanicIfError(err)

	fmt.Println(res)
}
func TestUpdateManyUnset(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateMany(context.Background(), bson.D{}, bson.D{
		{"$unset", bson.D{{"salah", ""}}},
	})
	helper.PanicIfError(err)

	fmt.Println(res)
}

func TestInsertManySalahLagi(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	docs := []interface{}{
		struct {
			Id    int    `bson:"_id,omitempty"`
			Name  string `bson:"name,omitempty"`
			Wrong string `bson:"wrong,omitempty"`
		}{
			Id:    9,
			Name:  "Ups Salah",
			Wrong: "Salah lagi",
		},
	}

	res, err := coll.InsertMany(context.Background(), docs)
	helper.PanicIfError(err)

	fmt.Println(res)
}
func TestReplaceOne(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.ReplaceOne(context.Background(), bson.D{{"_id", 9}}, model.Product{
		Name:     "Adidas Pulseboost HD Running Shoes Sepatu lari Pria",
		Price:    1100000,
		Category: "shoes",
		Tags: []string{
			"adidas", "shoes", "running",
		},
	})
	helper.PanicIfError(err)

	fmt.Println(res)
}

func TestUpdateCoba(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateOne(context.Background(), bson.D{{"_id", 1}}, bson.D{{
		"$set", model.Product{
			Name: "Indomie Ayam Bawang",
		},
	}})
	helper.PanicIfError(err)
	fmt.Println(res)
}
