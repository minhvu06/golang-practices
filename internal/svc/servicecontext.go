// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"context"
	"time"
	"user-api/internal/config"
	"user-api/internal/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceContext struct {
	Config  config.Config
	UserMod *model.UserModel

	mongoClient *mongo.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(c.Mongo.Uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:      c,
		UserMod:     model.NewUserModel(client, c.Mongo.Database),
		mongoClient: client,
	}
}

func (s *ServiceContext) MongoClient() *mongo.Client {
	return s.mongoClient
}
