package service

import (
	"context"
	"fmt"

	"github.com/TonyChinwe/libraryapp/utils/helpers"

	"github.com/TonyChinwe/libraryapp/pkg/books"
	"github.com/TonyChinwe/libraryapp/pkg/books/model"
	"github.com/TonyChinwe/libraryapp/pkg/library"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookService struct {
	bookRepo   books.BookRepositoryImpl
	LibService library.LibraryServiceImpl
}

func NewBookService(bookRepo books.BookRepositoryImpl, libservice library.LibraryServiceImpl) books.BookServiceImpl {
	return BookService{bookRepo: bookRepo, LibService: libservice}
}

func (service BookService) CreateBook(book model.Book, ctx context.Context) (model.Book, error) {
	result, err := service.bookRepo.CreateBook(book, ctx)
	if err != nil {
		return book, err
	}
	id := result.(primitive.ObjectID).Hex()
	fmt.Println("inserted id ", id)
	return service.GetBook(id, ctx)
}

func (service BookService) GetBook(id string, ctx context.Context) (model.Book, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}

	getBook, err := service.bookRepo.GetBook(filter, ctx)
	if err != nil {
		return getBook, err
	}
	fmt.Println("Record Found: ", getBook)
	return getBook, nil
}

// func (service BookService) GetAUser(username string, ctx context.Context) (userModel.User, error) {
// 	filter := bson.M{"username": username}
// 	user, err := service.bookRepo.GetAUser(filter, ctx)
// 	if err != nil {
// 		return user, err
// 	}
// 	fmt.Println("Record Found: ", user)
// 	return user, nil
// }

func (service BookService) DeleteBook(id string, ctx context.Context) (model.BookDeleted, error) {
	result := model.BookDeleted{
		DeletedCount: 0,
	}

	_id, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	if err != nil {
		return result, err
	}

	_, err = service.GetBook(id, ctx)
	if err != nil {
		return result, err
	}
	res, err := service.bookRepo.DeleteBook(filter, ctx)
	if err != nil {
		return result, err
	}
	result.DeletedCount = res
	return result, nil
}

func (service BookService) UpdateBook(id string, book model.Book, ctx context.Context) (model.BookUpdated, error) {
	result := model.BookUpdated{
		ModifiedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	filter := bson.M{"_id": _id}

	_, err = service.GetBook(id, ctx)
	if err != nil {
		return result, err
	}

	res, err := service.bookRepo.UpdateBook(filter, book, ctx)
	if err != nil {
		return result, err
	}

	result.ModifiedCount = 1
	result.Result = res
	fmt.Println("Record updated: ", result)
	return result, nil

}

func (service BookService) GetAllBooksInLibrary(library string, ctx context.Context) ([]model.Book, error) {
	filter := bson.M{"library": library}
	var books []model.Book
	cursor, err := service.bookRepo.GetAllBooksInLibrary(filter, ctx)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &books)
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, helpers.NewError("no result found")
	}
	return books, nil
}

func (service BookService) SearchBook(filter interface{}, ctx context.Context) ([]model.Book, error) {
	if filter == nil {
		filter = bson.M{}
	}
	var books []model.Book
	cursor, err := service.bookRepo.SearchBook(filter, ctx)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &books)
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, helpers.NewError("no result found")
	}
	return books, nil
}
