package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var C *mongo.Database

func Connect(host, db, user, password string) error {
	opt := options.Client()

	if user != "" && password != "" {
		opt.SetAuth(options.Credential{
			Username: user,
			Password: password,
		})
	}

	if host != "" {
		opt.ApplyURI(host)
	}

	client, err := mongo.NewClient(opt)
	if err != nil {
		return err
	}

	if err := client.Connect(context.Background()); err != nil {
		return err
	}

	C = client.Database(db)

	return nil
}
