package golang_mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"golang-mongodb/helper"
	"golang-mongodb/model"
	"testing"
)

func TestInsertManyProducts2(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")
	docs := []interface{}{
		model.Product{
			Id:       6,
			Name:     "Logitech M235 Wireless Mouse",
			Price:    175000,
			Category: "laptop",
			Tags: []string{
				"logitech", "mouse", "accessories",
			},
		},
		model.Product{
			Id:       7,
			Name:     "Havit Cooler Pad Gaming 5Fan Blue led F2082",
			Price:    200000,
			Category: "laptop",
			Tags: []string{
				"cooler", "laptop", "accessories", "fan",
			},
		},
		model.Product{
			Id:       8,
			Name:     "Samsung LC24F390FHEXXD Curved Monitor ",
			Price:    1750000,
			Category: "computer",
			Tags: []string{
				"samsung", "monitor", "computer",
			},
		},
	}
	res, err := coll.InsertMany(context.Background(), docs)
	helper.PanicIfError(err)
	fmt.Println(res)
}

func TestFindArrayAll(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	var model []model.Product
	cur, err := coll.Find(context.Background(), bson.D{{
		"tags", bson.D{{
			"$all", []string{"samsung", "monitor"},
		}},
	}})
	cur.All(context.Background(), &model)
	helper.PanicIfError(err)
	fmt.Println(model)
}
func TestFindArrayElemMatch(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	var model []model.Product
	cur, err := coll.Find(context.Background(), bson.D{{
		"tags", bson.D{{
			"$elemMatch", bson.D{
				{
					"$in", []string{"samsung", "logitech"},
				},
			},
		}},
	}})
	cur.All(context.Background(), &model)
	helper.PanicIfError(err)
	fmt.Println(model)
}
func TestFindArraySize(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	var model []model.Product
	cur, err := coll.Find(context.Background(), bson.D{{
		"tags", bson.D{{
			"$size", 3,
		}},
	}})
	cur.All(context.Background(), &model)
	helper.PanicIfError(err)
	fmt.Println(model)
}
