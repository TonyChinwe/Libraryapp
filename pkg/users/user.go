package users

import (
	"context"

	"github.com/TonyChinwe/libraryapp/pkg/users/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl interface {
	CreateUser(user model.User, ctx context.Context) (interface{}, error)
	// CreateUserAuthor(user model.User, ctx context.Context) (interface{}, error)
	// CreateUserAdmin(user model.User, ctx context.Context) (interface{}, error)

	AuthenticateUser(filter primitive.M, ctx context.Context) (model.User, error)

	GetUser(id primitive.M, ctx context.Context) (model.User, error)
	GetUserByUsername(id primitive.M, ctx context.Context) (model.User, error)
	DeleteUser(filter primitive.M, ctx context.Context) (int64, error)
	UpdateUser(filter primitive.M, book model.User, ctx context.Context) (model.User, error)
	GetAllUsers(ctx context.Context) (*mongo.Cursor, error)
	SearchUser(filter interface{}, ctx context.Context) (*mongo.Cursor, error)
}

type UserServiceImpl interface {
	CreateUser(user model.User, ctx context.Context) (model.User, error)
	// CreateUserAuthor(user model.User, ctx context.Context) (model.User, error)
	// CreateUserAdmin(user model.User, ctx context.Context) (model.User, error)

	AuthenticateUser(credentials *model.LoginReq, ctx context.Context) (*model.User, error)

	GetUser(id string, ctx context.Context) (model.User, error)
	GetUserByUsername(id string, ctx context.Context) (model.User, error)
	DeleteUser(id string, ctx context.Context) (model.UserDeleted, error)
	UpdateUser(id string, user model.User, ctx context.Context) (model.UserUpdated, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
	SearchUser(filter interface{}, ctx context.Context) ([]model.User, error)
}
