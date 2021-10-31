package golang_mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"golang-mongodb/helper"
	"testing"
)

func TestRenameFiled(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("customers")

	res, err := coll.UpdateMany(context.Background(), bson.D{}, bson.D{{
		"$rename", bson.D{{
			"name", "full_name",
		}},
	}})
	helper.PanicIfError(err)
	fmt.Println(res)
}
func TestIncrement(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateMany(context.Background(), bson.D{}, bson.D{{
		"$inc", bson.D{{
			"stock", 10,
		}},
	}})
	helper.PanicIfError(err)
	fmt.Println(res)
}
func TestCurrentDate(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateMany(context.Background(), bson.D{}, bson.D{{
		"$currentDate", bson.D{{
			"updated_at", bson.D{{
				"$type", "date",
			}},
		}},
	}})
	helper.PanicIfError(err)
	fmt.Println(res)
}
