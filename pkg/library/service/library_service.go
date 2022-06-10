package service

import (
	"context"
	"fmt"

	"github.com/TonyChinwe/libraryapp/utils/helpers"

	bookModel "github.com/TonyChinwe/libraryapp/pkg/books/model"
	"github.com/TonyChinwe/libraryapp/pkg/library"
	"github.com/TonyChinwe/libraryapp/pkg/library/model"

	userModel "github.com/TonyChinwe/libraryapp/pkg/users/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LibraryService struct {
	libraryRepo library.LibraryyRepositoryImpl
}

func NewLibraryService(libRepo library.LibraryyRepositoryImpl) library.LibraryServiceImpl {
	return LibraryService{libraryRepo: libRepo}
}

func (service LibraryService) GetAUser(username string, ctx context.Context) (userModel.User, error) {
	filter := bson.M{"username": username}
	user, err := service.libraryRepo.GetAUser(filter, ctx)
	if err != nil {
		return user, err
	}
	fmt.Println("Record Found: ", user)
	return user, nil
}

func (service LibraryService) GetLibraryByname(name string, ctx context.Context) (model.Library, error) {

	filter := bson.M{"name": name}
	lib, err := service.libraryRepo.GetLibraryByName(filter, ctx)

	if err != nil {
		return lib, err
	}
	return lib, nil
}

func (service LibraryService) CreateLibrary(library model.Library, ctx context.Context) (model.Library, error) {

	result, err := service.libraryRepo.CreateLibrary(library, ctx)
	if err != nil {
		return library, err
	}
	id := result.(primitive.ObjectID).Hex()
	fmt.Println("inserted id ", id)
	return service.GetLibrary(id, ctx)
}

func (service LibraryService) GetLibrary(id string, ctx context.Context) (model.Library, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}

	getLibrary, err := service.libraryRepo.GetLibrary(filter, ctx)
	if err != nil {
		return getLibrary, err
	}
	fmt.Println("Record Found: ", getLibrary)
	return getLibrary, nil
}
func (service LibraryService) DeleteLibrary(id string, ctx context.Context) (model.LibraryDeleted, error) {
	result := model.LibraryDeleted{
		DeletedCount: 0,
	}

	_id, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	if err != nil {
		return result, err
	}

	_, err = service.GetLibrary(id, ctx)
	if err != nil {
		return result, err
	}
	res, err := service.libraryRepo.DeleteLibrary(filter, ctx)
	if err != nil {
		return result, err
	}
	result.DeletedCount = res
	return result, nil
}

func (service LibraryService) UpdateLibrary(id string, library model.Library, ctx context.Context) (model.LibraryUpdated, error) {
	result := model.LibraryUpdated{
		ModifiedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	filter := bson.M{"_id": _id}

	_, err = service.GetLibrary(id, ctx)
	if err != nil {
		return result, err
	}

	res, err := service.libraryRepo.UpdateLibrary(filter, library, ctx)
	if err != nil {
		return result, err
	}

	result.ModifiedCount = 1
	result.Result = res
	fmt.Println("Record updated: ", result)
	return result, nil

}

func (service LibraryService) GetAllLibrary(ctx context.Context) ([]model.Library, error) {
	var libs []model.Library
	cursor, err := service.libraryRepo.GetAllLibrary(ctx)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &libs)
	if err != nil {
		return nil, err
	}
	if len(libs) == 0 {
		return nil, helpers.NewError("no result found")
	}
	return libs, nil
}

func (service LibraryService) GetAllBooksFromLibrary(id string, ctx context.Context) ([]bookModel.Book, error) {
	var buks []bookModel.Book
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}

	cursor, err := service.libraryRepo.GetAllBooksFromLibrary(filter, ctx)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &buks)
	if err != nil {
		return nil, err
	}
	if len(buks) == 0 {
		return nil, helpers.NewError("no result found")
	}
	return buks, nil
}

func (service LibraryService) SearchLibrary(filter interface{}, ctx context.Context) ([]model.Library, error) {
	if filter == nil {
		filter = bson.M{}
	}
	var libs []model.Library
	cursor, err := service.libraryRepo.SearchLibrary(filter, ctx)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &libs)
	if err != nil {
		return nil, err
	}
	if len(libs) == 0 {
		return nil, helpers.NewError("no result found")
	}
	return libs, nil
}
