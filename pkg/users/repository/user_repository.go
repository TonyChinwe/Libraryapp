package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	dbs "github.com/TonyChinwe/libraryapp/config/db"
	"github.com/TonyChinwe/libraryapp/pkg/users"
	"github.com/TonyChinwe/libraryapp/pkg/users/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepo struct {
	conn *dbs.MongoDB
}

func NewUserRepo(conn *dbs.MongoDB) users.UserRepositoryImpl {
	return UserRepo{conn: conn}
}

func (repo UserRepo) AuthenticateUser(filter primitive.M, ctx context.Context) (model.User, error) {
	return repo.GetUser(filter, ctx)
}

func (repo UserRepo) CreateUser(user model.User, ctx context.Context) (interface{}, error) {
	userColl := repo.conn.UserCollections()
	result, err := userColl.InsertOne(ctx, user)
	return result.InsertedID, err
}

// func (repo UserRepo) CreateUserAuthor(user model.User, ctx context.Context) (interface{}, error) {
// 	userColl := repo.conn.BookCollections()
// 	result, err := userColl.InsertOne(ctx, user)
// 	return result.InsertedID, err
// }

// func (repo UserRepo) CreateUserAdmin(user model.User, ctx context.Context) (interface{}, error) {
// 	userColl := repo.conn.BookCollections()
// 	result, err := userColl.InsertOne(ctx, user)
// 	return result.InsertedID, err
// }

func (repo UserRepo) GetUser(filter primitive.M, ctx context.Context) (model.User, error) {
	var us model.User
	userColl := repo.conn.UserCollections()
	err := userColl.FindOne(ctx, filter).Decode(&us)
	fmt.Println("guy " + us.Username)
	return us, err

}

func (repo UserRepo) GetUserByUsername(filter primitive.M, ctx context.Context) (model.User, error) {
	var us model.User
	userColl := repo.conn.UserCollections()
	err := userColl.FindOne(ctx, filter).Decode(&us)
	fmt.Println("guy " + us.Username)
	return us, err

}

func (repo UserRepo) DeleteUser(filter primitive.M, ctx context.Context) (int64, error) {
	userColl := repo.conn.UserCollections()
	result, err := userColl.DeleteOne(ctx, filter)
	return result.DeletedCount, err
}
func (repo UserRepo) UpdateUser(filter primitive.M, user model.User, ctx context.Context) (model.User, error) {
	userColl := repo.conn.UserCollections()
	err := userColl.FindOneAndUpdate(ctx, filter, bson.M{"$set": user}, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&user)
	return user, err
}
func (repo UserRepo) GetAllUsers(ctx context.Context) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetLimit(100)
	filter := bson.D{{}}
	userColl := repo.conn.UserCollections()
	cursor, err := userColl.Find(ctx, filter, findOptions)
	return cursor, err
}

func (repo UserRepo) SearchUser(filter interface{}, ctx context.Context) (*mongo.Cursor, error) {
	findOptions := options.Find()
	findOptions.SetLimit(100)
	userColl := repo.conn.UserCollections()
	cursor, err := userColl.Find(ctx, filter, findOptions)
	return cursor, err
}
