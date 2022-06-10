package model

import (
	libModel "github.com/TonyChinwe/libraryapp/pkg/library/model"
	"github.com/TonyChinwe/libraryapp/pkg/users/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title   string             `bson:"title,omitempty" json:"title,omitempty"`
	ISBN    string             `bson:"isbn,omitempty" json:"isbn,omitempty"`
	Author  *model.User        `bson:"author,omitempty" json:"author,omitempty"`
	Library *libModel.Library  `bson:"library,omitempty" json:"library,omitempty"`
	LibName string             `bson:"libname,omitempty" json:"libname,omitempty"`
}

type BookUpdated struct {
	ModifiedCount int64 `json:"modifiedcount"`
	Result        Book
}

type BookDeleted struct {
	DeletedCount int64 `json:"deletedcount"`
}
