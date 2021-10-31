package golang_mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-mongodb/helper"
	"golang-mongodb/model"
	"testing"
)

func TestGetConnection(t *testing.T) {
	GetConnection()
}

func TestInsertOne(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("customers")
	res, err := coll.InsertOne(context.Background(), model.Customer{
		Id:   "khannedy",
		FullName: "Eko Kurniawan Khannedy",
	})
	helper.PanicIfError(err)
	fmt.Println(res)
}

func TestInsertMany(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")
	docs := []interface{}{
		model.Product{
			Id:    1,
			Name:  "indomie Ayam Bawang",
			Price: 2000,
		},
		model.Product{
			Id:    2,
			Name:  "Mie Sedap",
			Price: 2000,
		},
	}
	// If true, no writes will be executed after one fails. The default value is true.
	opts := options.InsertMany().SetOrdered(true)
	res, err := coll.InsertMany(context.Background(), docs, opts)
	helper.PanicIfError(err)
	fmt.Println(res)
}

func TestInsertOne2(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("orders")
	res, err := coll.InsertOne(context.Background(), model.Order{
		// jika field id kosong secara default id akan diisi dengan objectid
		//Id:    primitive.ObjectID{},
		Total: 8000,
		Items: []model.Item{
			{
				ProductId: 1,
				Price:     2000,
				Quantity:  2,
			},
			{
				ProductId: 2,
				Price:     2000,
				Quantity:  2,
			},
		},
	})
	helper.PanicIfError(err)
	fmt.Println(res)
}
