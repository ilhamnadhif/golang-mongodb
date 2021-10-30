package golang_mongodb

import (
	"golang-mongodb/helper"
	"golang-mongodb/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestFindOneEq(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("customers")

	var customer model.Customer
	err := coll.FindOne(context.Background(), bson.D{{
		"_id", bson.D{{
			"$eq", "khannedy",
		}},
	}}).Decode(&customer)
	helper.PanicIfError(err)
	fmt.Println(customer)
}
func TestFindOneGt(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	var model model.Product
	err := coll.FindOne(context.Background(), bson.D{{
		"price", bson.D{{
			"$gt", 2500, // grather than / lebih dari
		}},
	}}).Decode(&model)
	helper.PanicIfError(err)
	fmt.Println(model)
}
func TestInsertManyProducts(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")
	docs := []interface{}{
		model.Product{
			Id:       3,
			Name:     "Pop Mie Rasa Bakso",
			Price:    2500,
			Category: "food",
		},
		model.Product{
			Id:       4,
			Name:     "Samsung Galaxy S9+",
			Price:    10000000,
			Category: "handphone",
		},
		model.Product{
			Id:       5,
			Name:     "Acer Precator XXI",
			Price:    25000000,
			Category: "laptop",
		},
	}
	//opts := options.InsertMany().SetOrdered(true)
	res, err := coll.InsertMany(context.Background(), docs)
	helper.PanicIfError(err)
	fmt.Println(res)
}

func TestFindIn(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	var model []model.Product
	cur, err := coll.Find(context.Background(), bson.D{{
		"category", bson.D{{
			"$in", []string{"handphone", "laptop"},
		}},
	},
		{"price", bson.D{
			{"$gt", 5000000},
		}},
	})
	cur.All(context.Background(), &model)
	helper.PanicIfError(err)
	fmt.Println(model)
}
