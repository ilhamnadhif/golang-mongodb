package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	Id   string `bson:"_id,omitempty"`
	Name string `bson:"name,omitempty"`
}

type Product struct {
	Id       int      `bson:"_id,omitempty"`
	Name     string   `bson:"name,omitempty"`
	Price    int      `bson:"price,omitempty"`
	Category string   `bson:"category,omitempty"`
	Tags     []string `bson:"tags,omitempty"`
}

type Item struct {
	ProductId int `bson:"product_id,omitempty"`
	Price     int `bson:"price,omitempty"`
	Quantity  int `bson:"quantity,omitempty"`
}

type Order struct {
	Id    primitive.ObjectID `bson:"_id,omitempty"`
	Total int                `bson:"total,omitempty"`
	Items []Item             `bson:"items,omitempty"`
}
