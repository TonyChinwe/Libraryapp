package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	dbs "github.com/TonyChinwe/libraryapp/config/db"
	"github.com/TonyChinwe/libraryapp/pkg/library"
	"github.com/TonyChinwe/libraryapp/pkg/library/model"
	userModel "github.com/TonyChinwe/libraryapp/pkg/users/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LibraryRepo struct {
	conn *dbs.MongoDB
}

func NewLibraryRepo(conn *dbs.MongoDB) library.LibraryyRepositoryImpl {
	return LibraryRepo{conn: conn}
}

func (repo LibraryRepo) GetLibraryByName(filter primitive.M, ctx context.Context) (model.Library, error) {
	var lib model.Library
	libColl := repo.conn.LibraryCollections()
	err := libColl.FindOne(ctx, filter).Decode(&lib)
	return lib, err

}

func (repo LibraryRepo) GetAUser(filter primitive.M, ctx context.Context) (userModel.User, error) {
	var user userModel.User
	userColl := repo.conn.UserCollections()
	err := userColl.FindOne(ctx, filter).Decode(&user)
	return user, err

}

func (repo LibraryRepo) CreateLibrary(library model.Library, ctx context.Context) (interface{}, error) {
	libColl := repo.conn.LibraryCollections()
	result, err := libColl.InsertOne(ctx, library)
	return result.InsertedID, err
}
func (repo LibraryRepo) GetLibrary(filter primitive.M, ctx context.Context) (model.Library, error) {
	var lib model.Library
	libColl := repo.conn.LibraryCollections()
	err := libColl.FindOne(ctx, filter).Decode(&lib)
	return lib, err

}
func (repo LibraryRepo) DeleteLibrary(filter primitive.M, ctx context.Context) (int64, error) {
	libColl := repo.conn.LibraryCollections()
	result, err := libColl.DeleteOne(ctx, filter)
	return result.DeletedCount, err
}
func (repo LibraryRepo) UpdateLibrary(filter primitive.M, library model.Library, ctx context.Context) (model.Library, error) {
	libColl := repo.conn.LibraryCollections()
	err := libColl.FindOneAndUpdate(ctx, filter, bson.M{"$set": library}, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&library)
	return library, err
}
func (repo LibraryRepo) GetAllLibrary(ctx context.Context) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetLimit(100)
	filter := bson.D{{}}
	libColl := repo.conn.LibraryCollections()
	cursor, err := libColl.Find(ctx, filter, findOptions)
	return cursor, err
}

func (repo LibraryRepo) GetAllBooksFromLibrary(filter primitive.M, ctx context.Context) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetLimit(100)
	// filter := bson.D{{}}
	libColl := repo.conn.LibraryCollections()
	cursor, err := libColl.Find(ctx, filter, findOptions)
	return cursor, err
}

func (repo LibraryRepo) SearchLibrary(filter interface{}, ctx context.Context) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetLimit(100)
	libColl := repo.conn.LibraryCollections()
	cursor, err := libColl.Find(ctx, filter, findOptions)
	return cursor, err
}
