package service

import (
	"context"
	"fmt"

	"github.com/TonyChinwe/libraryapp/utils/helpers"

	"github.com/TonyChinwe/libraryapp/pkg/users"
	"github.com/TonyChinwe/libraryapp/pkg/users/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	userRepo users.UserRepositoryImpl
}

func NewUserService(userRepo users.UserRepositoryImpl) users.UserServiceImpl {
	return UserService{userRepo: userRepo}
}

func (service UserService) AuthenticateUser(credentials *model.LoginReq, ctx context.Context) (*model.User, error) {
	filter := bson.M{"username": credentials.Username}
	authenticateUser, err := service.userRepo.AuthenticateUser(filter, ctx)

	if err != nil {
		return nil, helpers.NewError("invalid username")
	}
	err = authenticateUser.ComparePassword(credentials.Password)
	if err != nil {
		fmt.Println("Incorrect guy mam")
		return nil, helpers.NewError("invalid password")
	}
	return &authenticateUser, nil
}

func (service UserService) CreateUser(user model.User, ctx context.Context) (model.User, error) {
	result, err := service.userRepo.CreateUser(user, ctx)
	if err != nil {
		return user, err
	}
	id := result.(primitive.ObjectID).Hex()
	fmt.Println("inserted id ", id)
	return service.GetUser(id, ctx)
}

func (service UserService) GetUser(id string, ctx context.Context) (model.User, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}

	getUser, err := service.userRepo.GetUser(filter, ctx)
	if err != nil {
		return getUser, err
	}
	fmt.Println("Record Found: ", getUser)
	return getUser, nil
}

func (service UserService) GetUserByUsername(username string, ctx context.Context) (model.User, error) {

	filter := bson.M{"username": username}
	authenticateUser, err := service.userRepo.AuthenticateUser(filter, ctx)

	if err != nil {
		return authenticateUser, err
	}
	fmt.Println("Record Found: ", authenticateUser)
	return authenticateUser, nil
}

func (service UserService) DeleteUser(id string, ctx context.Context) (model.UserDeleted, error) {
	result := model.UserDeleted{
		DeletedCount: 0,
	}

	_id, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	if err != nil {
		return result, err
	}

	_, err = service.GetUser(id, ctx)
	if err != nil {
		return result, err
	}
	res, err := service.userRepo.DeleteUser(filter, ctx)
	if err != nil {
		return result, err
	}
	result.DeletedCount = res
	return result, nil
}

func (service UserService) UpdateUser(id string, user model.User, ctx context.Context) (model.UserUpdated, error) {
	result := model.UserUpdated{
		ModifiedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	filter := bson.M{"_id": _id}

	_, err = service.GetUser(id, ctx)
	if err != nil {
		return result, err
	}

	res, err := service.userRepo.UpdateUser(filter, user, ctx)
	if err != nil {
		return result, err
	}

	result.ModifiedCount = 1
	result.Result = res
	fmt.Println("Record updated: ", result)
	return result, nil

}

func (service UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	cursor, err := service.userRepo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, helpers.NewError("no result found")
	}
	return users, nil
}

func (service UserService) SearchUser(filter interface{}, ctx context.Context) ([]model.User, error) {
	if filter == nil {
		filter = bson.M{}
	}
	var users []model.User
	cursor, err := service.userRepo.SearchUser(filter, ctx)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, helpers.NewError("no result found")
	}
	return users, nil
}
