package handlers

import (
	"context"
	"fmt"
	"ipumpkin/config"

	// "ipumpkin/config"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func Fyhtest() int {
	fmt.Println(111)
	return 111
}

func FyhMongo(cfg config.Provider, mongodb *mongo.Database, logger *logrus.Logger) {
	collection := mongodb.Collection(cfg.GetString("mongodb.docker"))
	collection.InsertOne(context.Background(), 111)
	logger.Info("")
}
