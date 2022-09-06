package handler

import (
	"context"
	"fmt"
	"ipumpkin/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB(cfg config.Provider) *mongo.Database {
	host := cfg.GetString("mongo.host")
	port := cfg.GetString("mongo.port")
	user := cfg.GetString("mongo.user")
	passwd := cfg.GetString("mongo.passwd")
	table := cfg.GetString("mongotable.table")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	url := fmt.Sprintf("mongodb://%v:%v@%v:%v", user, passwd, host, port)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(url))
	mongodb := client.Database(table)
	return mongodb
}
