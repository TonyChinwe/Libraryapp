package library

import (
	"context"

	book "github.com/TonyChinwe/libraryapp/pkg/books/model"
	"github.com/TonyChinwe/libraryapp/pkg/library/model"
	userModel "github.com/TonyChinwe/libraryapp/pkg/users/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LibraryyRepositoryImpl interface {
	GetAUser(id primitive.M, ctx context.Context) (userModel.User, error)

	CreateLibrary(library model.Library, ctx context.Context) (interface{}, error)
	GetLibrary(id primitive.M, ctx context.Context) (model.Library, error)
	DeleteLibrary(filter primitive.M, ctx context.Context) (int64, error)
	UpdateLibrary(filter primitive.M, library model.Library, ctx context.Context) (model.Library, error)
	GetAllLibrary(ctx context.Context) (*mongo.Cursor, error)
	SearchLibrary(filter interface{}, ctx context.Context) (*mongo.Cursor, error)
	GetAllBooksFromLibrary(libId primitive.M, ctx context.Context) (*mongo.Cursor, error)
	GetLibraryByName(id primitive.M, ctx context.Context) (model.Library, error)
}

type LibraryServiceImpl interface {
	GetAUser(id string, ctx context.Context) (userModel.User, error)
	GetLibraryByname(id string, ctx context.Context) (model.Library, error)

	CreateLibrary(library model.Library, ctx context.Context) (model.Library, error)
	GetLibrary(id string, ctx context.Context) (model.Library, error)
	DeleteLibrary(id string, ctx context.Context) (model.LibraryDeleted, error)
	UpdateLibrary(id string, library model.Library, ctx context.Context) (model.LibraryUpdated, error)
	GetAllLibrary(ctx context.Context) ([]model.Library, error)
	GetAllBooksFromLibrary(id string, ctx context.Context) ([]book.Book, error)
	SearchLibrary(filter interface{}, ctx context.Context) ([]model.Library, error)
}
