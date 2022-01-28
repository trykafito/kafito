package post

import (
	"context"
	"time"

	"github.com/trykafito/kafito/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Post struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	CreatedBy   primitive.ObjectID `bson:"created_by" json:"created_by"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Tags        []string           `bson:"tags" json:"tags"`
	Thumbnail   primitive.ObjectID `bson:"thumbnail" json:"thumbnail"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

func collection() *mongo.Collection {
	return database.C.Collection("posts")
}

func Count(filter bson.M) int {
	count, _ := collection().CountDocuments(context.Background(), filter)
	return int(count)
}

func Find(filter bson.M, page, limit int, sorts ...bson.E) ([]Post, error) {
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

	posts := []Post{}
	for cursor.Next(context.Background()) {
		p := new(Post)
		if err := cursor.Decode(p); err != nil {
			return nil, err
		}

		posts = append(posts, *p)
	}

	return posts, nil
}

func FindOne(filter bson.M) (*Post, error) {
	p := new(Post)
	if err := collection().FindOne(context.Background(), filter).Decode(p); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Post) Insert() error {
	p.ID = primitive.NewObjectID()
	p.CreatedAt = time.Now()

	_, err := collection().InsertOne(context.Background(), p)
	return err
}

func (p *Post) Update() error {
	_, err := collection().UpdateOne(context.Background(), bson.M{"_id": p.ID}, bson.M{"$set": p})
	return err
}

func (p *Post) Save() error {
	if p.ID.IsZero() {
		return p.Insert()
	}

	return p.Update()
}

func (p *Post) Delete() error {
	_, err := collection().DeleteOne(context.Background(), bson.M{"_id": p.ID})
	return err
}
