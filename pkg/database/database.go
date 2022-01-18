package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
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

func Bson(d interface{}) bson.M {
	val, _ := bson.Marshal(d)
	data := new(bson.M)
	bson.Unmarshal(val, data)

	return *data
}
