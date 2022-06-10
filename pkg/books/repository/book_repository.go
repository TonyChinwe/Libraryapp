package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	dbs "github.com/TonyChinwe/libraryapp/config/db"
	"github.com/TonyChinwe/libraryapp/pkg/books"
	"github.com/TonyChinwe/libraryapp/pkg/books/model"
	libModel "github.com/TonyChinwe/libraryapp/pkg/library/model"
	userModel "github.com/TonyChinwe/libraryapp/pkg/users/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BookRepo struct {
	conn *dbs.MongoDB
}

func NewBookRepo(conn *dbs.MongoDB) books.BookRepositoryImpl {
	return BookRepo{conn: conn}
}

func (repo BookRepo) CreateBook(book model.Book, ctx context.Context) (interface{}, error) {
	bookColl := repo.conn.BookCollections()
	result, err := bookColl.InsertOne(ctx, book)
	return result.InsertedID, err
}

func (repo BookRepo) GetBook(filter primitive.M, ctx context.Context) (model.Book, error) {
	var buk model.Book
	bookColl := repo.conn.BookCollections()
	err := bookColl.FindOne(ctx, filter).Decode(&buk)
	return buk, err

}

func (repo BookRepo) GetAUser(filter primitive.M, ctx context.Context) (userModel.User, error) {
	var user userModel.User
	userColl := repo.conn.UserCollections()
	err := userColl.FindOne(ctx, filter).Decode(&user)
	return user, err

}
func (repo BookRepo) DeleteBook(filter primitive.M, ctx context.Context) (int64, error) {
	bookColl := repo.conn.BookCollections()
	result, err := bookColl.DeleteOne(ctx, filter)
	return result.DeletedCount, err
}
func (repo BookRepo) UpdateBook(filter primitive.M, book model.Book, ctx context.Context) (model.Book, error) {
	bookColl := repo.conn.BookCollections()
	err := bookColl.FindOneAndUpdate(ctx, filter, bson.M{"$set": book}, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&book)
	return book, err
}
func (repo BookRepo) GetAllBooksInLibrary(library primitive.M, ctx context.Context) (*mongo.Cursor, error) {
	var libry libModel.Library
	libColl := repo.conn.LibraryCollections()
	libColl.FindOne(ctx, library).Decode(&libry)
	findOptions := options.Find()
	findOptions.SetLimit(100)
	//  filter := bson.D{library}
	bookColl := repo.conn.BookCollections()
	cursor, err := bookColl.Find(ctx, libry, findOptions)

	return cursor, err
}

func (repo BookRepo) SearchBook(filter interface{}, ctx context.Context) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetLimit(100)
	bookColl := repo.conn.BookCollections()
	cursor, err := bookColl.Find(ctx, filter, findOptions)
	return cursor, err
}
