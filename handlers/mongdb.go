package handlers

import (
	"context"
	"fmt"
	"ipumpkin/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitDB(cfg config.Provider) *mongo.Database {
	host := cfg.GetString("mongo.host")
	port := cfg.GetString("mongo.port")
	user := cfg.GetString("mongo.user")
	passwd := cfg.GetString("mongo.passwd")
	db := cfg.GetString("mongo.dbname")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// option := options.Client().ApplyURI("mongodb://localhost:27017")
	// client, err := mongo.Connect(context.TODO(), option)
	defer cancel()
	url := fmt.Sprintf("mongodb://%v:%v@%v:%v", user, passwd, host, port)
	clientOptions := options.Client().ApplyURI(url)
	// 连接 MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(fmt.Sprintf("db conn err: %v", err))
	}

	err = client.Ping(ctx, readpref.Nearest())
	if err != nil {
		panic(fmt.Sprintf("db ping err: %v", err))
	}

	return client.Database(db)
}
