package infrastructure

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongo(config config.Config) (*mongo.Database, error) {
	uri := fmt.Sprintf("mongodb://%s", config.Mongo.Address)
	auth := options.Credential{
		Username:   config.Mongo.Username,
		Password:   config.Mongo.Password,
		AuthSource: config.Mongo.DatabaseName,
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri).SetAuth(auth).SetDirect(true))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return nil, errors.WithStack(err)
	}

	db := client.Database(config.Mongo.DatabaseName)

	return db, nil
}
