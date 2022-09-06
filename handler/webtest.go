package handler

import (
	"context"
	"fmt"
	"ipumpkin/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Fyhtest() int {
	fmt.Println(111)
	return 111
}

func FyhMongo(c context.Context, cfg config.Provider, mongodb *mongo.Database) {
	collection := mongodb.Collection(cfg.GetString("mongodb.test"))
	collection.InsertOne(context.Background(), bson.M{"test": 1111})
}
