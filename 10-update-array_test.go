package golang_mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-mongodb/helper"
	"testing"
)

func TestUpdateProductRating(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateMany(context.Background(), bson.D{}, bson.D{{
		"$set", bson.D{{
			"ratings", []int{90, 80, 70},
		}},
	}})
	helper.PanicIfError(err)
	fmt.Println(res)
}
func TestUpdateArrayDataPertama(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	// update sesuai query nya dan data yang pertama
	res, err := coll.UpdateMany(context.Background(), bson.D{{"ratings", 90}}, bson.D{{
		"$set", bson.D{{
			"ratings.$", 100,
		}},
	}})
	helper.PanicIfError(err)
	fmt.Println(res)
}
func TestUpdateArrayAllData(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateMany(context.Background(), bson.D{}, bson.D{{
		"$set", bson.D{{
			"ratings.$[]", 100,
		}},
	}})
	helper.PanicIfError(err)
	fmt.Println(res)
}
func TestArrayFilter(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateMany(context.Background(), bson.D{}, bson.D{
		{
			"$set", bson.D{{"ratings.$[element]", 100}},
		},
	}, options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: bson.A{
			bson.D{{
				"element", bson.D{{
					"$gte", 80,
				}},
			}},
		},
	}))
	helper.PanicIfError(err)
	fmt.Println(res)
}
func TestUpdateByIndexArray(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateMany(context.Background(), bson.D{}, bson.D{{
		"$set", bson.D{
			{
				"ratings.0", 50,
			},
			{
				"ratings.1", 60,
			},
		},
	}})
	helper.PanicIfError(err)
	fmt.Println(res)
}
func TestAddToSet(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateMany(context.Background(), bson.D{
		{"_id", 1},
	}, bson.D{{
		"$addToSet", bson.D{{
			"tags", "popular",
		}},
	}})
	helper.PanicIfError(err)
	fmt.Println(res)
}
func TestArrayPop(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateMany(context.Background(), bson.D{}, bson.D{{
		"$pop", bson.D{{
			"ratings", -1, // Menghapus element pertama (-1) atau terakhir (1) array
		}},
	}})
	helper.PanicIfError(err)
	fmt.Println(res)
}
func TestArrayPull(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateMany(context.Background(), bson.D{}, bson.D{{
		"$pull", bson.D{{
			"ratings", bson.D{{
				"$gte", 80,
			}},
		}},
	}})
	helper.PanicIfError(err)
	fmt.Println(res)
}
func TestArrayPush(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateMany(context.Background(), bson.D{}, bson.D{{
		"$push", bson.D{{
			"ratings", 100,
		}},
	}})
	helper.PanicIfError(err)
	fmt.Println(res)
}
func TestArrayPullAll(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateMany(context.Background(), bson.D{}, bson.D{{
		"$pullAll", bson.D{{
			"ratings", bson.A{100},
		}},
	}})
	helper.PanicIfError(err)
	fmt.Println(res)
}
func TestArrayEachPush(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateMany(context.Background(), bson.D{}, bson.D{{
		"$push", bson.D{{
			"ratings", bson.D{{
				"$each", []int{100, 200, 300},
			}},
		}},
	}})
	helper.PanicIfError(err)
	fmt.Println(res)
}
func TestArrayEachAddToSet(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateMany(context.Background(), bson.D{}, bson.D{{
		"$addToSet", bson.D{{
			"tags", bson.D{{
				"$each", []string{"trending", "popular"},
			}},
		}},
	}})
	helper.PanicIfError(err)
	fmt.Println(res)
}
func TestArrayPosition(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateMany(context.Background(), bson.D{}, bson.D{{
		"$push", bson.D{{
			"tags", bson.D{
				{
					"$each", []string{"hot"},
				},
				{
					"$position", 1,
				},
			},
		}},
	}})
	helper.PanicIfError(err)
	fmt.Println(res)
}
func TestArraySlice(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateMany(context.Background(), bson.D{}, bson.D{{
		"$push", bson.D{{
			"ratings", bson.D{
				{
					"$each", []int{100, 200, 300, 400, 500},
				},
				{
					"$slice", -4,
				},
			},
		}},
	}})
	helper.PanicIfError(err)
	fmt.Println(res)
}
func TestArraySort(t *testing.T) {
	client := GetConnection()
	defer CloseConnection(client)
	database := client.Database("belajar-mongodb")
	coll := database.Collection("products")

	res, err := coll.UpdateMany(context.Background(), bson.D{}, bson.D{{
		"$push", bson.D{{
			"ratings", bson.D{
				{
					"$each", []int{100, 200, 300, 400, 500},
				},
				{
					"$sort", -1,
				},
			},
		}},
	}})
	helper.PanicIfError(err)
	fmt.Println(res)
}
