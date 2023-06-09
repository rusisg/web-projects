package store

import (
	"context"

	"github.com/rusisg/e-commerce/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductStore struct {
	db   *mongo.Database
	coll string
}

func NewMongoProductStore(db *mongo.Database) *MongoProductStore {
	return &MongoProductStore{
		db:   db,
		coll: "products",
	}
}

func (s *MongoProductStore) Insert(ctx context.Context, p *types.Product) error {
	_, err := s.db.Collection(s.coll).InsertOne(ctx, p)

	return err
}

func (s *MongoProductStore) GetByID(ctx context.Context, id string) (*types.Product, error) {
	var (
		objID, _ = primitive.ObjectIDFromHex(id)
		res      = s.db.Collection(s.coll).FindOne(ctx, bson.M{"_id": objID})
		p        = &types.Product{}
		err      = res.Decode(p)
	)
	return p, err
}
