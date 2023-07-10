package mongo

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client *mongo.Client
	db     *mongo.Database
	Exams  *mongo.Collection
}

func GetConnectionStringFromEnvFile(filePath string) string {
	err := godotenv.Load(filePath)
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)

		panic(err)
	}
	return os.Getenv("MONGO_CONNECTION_STRING")
}

func New(uri string, dbName string) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	examsCollection := db.Collection("exams")
	return &Database{client: client, db: db, Exams: examsCollection}, nil
}
