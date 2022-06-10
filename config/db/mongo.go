package db

import (
	"context"
	"log"
	"time"

	"github.com/TonyChinwe/libraryapp/config/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Database *mongo.Database
}

func NewMongoDBConnection() *MongoDB {
	dbConn := &MongoDB{Database: connect()}
	return dbConn
}

func connect() *mongo.Database {
	dbName := env.GetEnvWithKey("DB_NAME")
	url := env.GetEnvWithKey("LOCAL_URL")
	clientOption := options.Client().ApplyURI(url)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)

	defer cancel()
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database", err)
	}
	return client.Database(dbName)
}

func (db MongoDB) BookCollections() *mongo.Collection {
	return db.Database.Collection("book")
}

func (db MongoDB) UserCollections() *mongo.Collection {
	return db.Database.Collection("user")
}

func (db MongoDB) AuthorCollections() *mongo.Collection {
	return db.Database.Collection("author")
}

func (db MongoDB) LibraryCollections() *mongo.Collection {
	return db.Database.Collection("library")
}
