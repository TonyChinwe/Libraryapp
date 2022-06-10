package books

import (
	"context"

	"github.com/TonyChinwe/libraryapp/pkg/books/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepositoryImpl interface {
	CreateBook(book model.Book, ctx context.Context) (interface{}, error)
	GetBook(id primitive.M, ctx context.Context) (model.Book, error)
	DeleteBook(filter primitive.M, ctx context.Context) (int64, error)
	UpdateBook(filter primitive.M, book model.Book, ctx context.Context) (model.Book, error)
	GetAllBooksInLibrary(filter primitive.M, ctx context.Context) (*mongo.Cursor, error)
	SearchBook(filter interface{}, ctx context.Context) (*mongo.Cursor, error)
}

type BookServiceImpl interface {
	CreateBook(book model.Book, ctx context.Context) (model.Book, error)
	GetBook(id string, ctx context.Context) (model.Book, error)
	DeleteBook(id string, ctx context.Context) (model.BookDeleted, error)
	UpdateBook(id string, book model.Book, ctx context.Context) (model.BookUpdated, error)
	GetAllBooksInLibrary(id string, ctx context.Context) ([]model.Book, error)
	SearchBook(filter interface{}, ctx context.Context) ([]model.Book, error)
}
