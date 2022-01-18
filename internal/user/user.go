package user

import (
	"context"
	"time"

	"github.com/trykafito/kafito/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Phone     string             `bson:"phone" json:"phone"`
	Password  string             `bson:"password" json:"-"`
	SuperUser bool               `bson:"super_user" json:"super_user"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

func collection() *mongo.Collection {
	return database.C.Collection("users")
}

func Count(filter bson.M) int {
	count, _ := collection().CountDocuments(context.Background(), filter)
	return int(count)
}

func Find(filter bson.M, page, limit int, sorts ...bson.E) ([]User, error) {
	opt := options.Find()
	opt.SetSort(sorts)

	if limit > 0 {
		opt.SetLimit(int64(limit))
	}

	if page > 1 {
		opt.SetSkip(int64((page - 1) * limit))
	}

	cursor, err := collection().Find(context.Background(), filter, opt)
	if err != nil {
		return nil, err
	}

	users := []User{}
	for cursor.Next(context.Background()) {
		u := new(User)
		if err := cursor.Decode(u); err != nil {
			return nil, err
		}

		users = append(users, *u)
	}

	return users, nil
}

func FindOne(filter bson.M) (*User, error) {
	u := new(User)
	if err := collection().FindOne(context.Background(), filter).Decode(u); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) Insert() error {
	u.ID = primitive.NewObjectID()
	u.CreatedAt = time.Now()

	_, err := collection().InsertOne(context.Background(), database.Bson(u))
	return err
}

func (u *User) Update() error {
	_, err := collection().UpdateOne(context.Background(), bson.M{"_id": u.ID}, bson.M{"$set": database.Bson(u)})
	return err
}

func (u *User) Save() error {
	if u.ID.IsZero() {
		return u.Insert()
	}

	return u.Update()
}

func (u *User) Delete() error {
	_, err := collection().DeleteOne(context.Background(), bson.M{"_id": u.ID})
	return err
}
