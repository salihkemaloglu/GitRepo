package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type Item struct {
	ID          bson.ObjectId `bson:"_id" json:"id" `
	Name        *string       `bson:"name" json:"name"`
	Value       *string       `bson:"value" json:"value"`
	Description string        `bson:"description" json:"description"`
	ItemId      string        `bson:"-" json:"-"`
	Count       int           `bson:"count" json:"count"`
}

func test() {
	fmt.Print("working")
}
