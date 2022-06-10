package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Library struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string             `bson:"name,omitempty" json:"name,omitempty"`
}

type LibraryUpdated struct {
	ModifiedCount int64 `json:"modifiedcount"`
	Result        Library
}

type LibraryDeleted struct {
	DeletedCount int64 `json:"deletedcount"`
}
