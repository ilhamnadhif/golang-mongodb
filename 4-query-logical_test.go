package golang_mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"golang-mongodb/helper"
	"golang-mongodb/model"
	"testing"
)

func TestFindAnd(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	var model []model.Product
	cur, err := coll.Find(context.Background(), bson.D{{
		"$and", bson.A{
			bson.D{{
				"category", bson.D{
					{
						"$in", []string{"laptop", "handphone"},
					},
				},
			}},
			bson.D{{
				"price", bson.D{
					{
						"$gt", 20000000,
					},
				},
			}},
		},
	}})
	cur.All(context.Background(), &model)
	helper.PanicIfError(err)
	fmt.Println(model)
}
func TestFindNotIn(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	var model []model.Product
	cur, err := coll.Find(context.Background(), bson.D{
		{
			"category", bson.D{
				{
					"$not", bson.D{
						{
							"$in", []string{"laptop", "handphone"},
						},
					},
				},
			},
		},
	})
	cur.All(context.Background(), &model)
	helper.PanicIfError(err)
	fmt.Println(model)
}
func TestFindBetween(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	var model []model.Product
	cur, err := coll.Find(context.Background(), bson.D{
		{
			"$and", bson.A{
				bson.D{
					{
						"price", bson.D{
							{
								"$gte", 10000000,
							},
							{
								"$lte", 20000000,
							},
						},
					},
				},
				bson.D{
					{
						"category", bson.D{{"$ne", "food"}},
					},
				},
			},
		},
	})
	cur.All(context.Background(), &model)
	helper.PanicIfError(err)
	fmt.Println(model)
}
