package dbManager

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Exam struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name         string             `json:"name" bson:"name"`
	IsInProgress bool               `json:"isInProgress" bson:"is_in_progress"`
	Challenger   string             `json:"challenger" bson:"challenger"`
	StartTime    int                `json:"startTime" bson:"start_time"`
	EndTime      int                `json:"endTime" bson:"end_time"`
}

func (db *Database) CreateExam(exam Exam) (*mongo.InsertOneResult, error) {
	result, err := db.Exams.InsertOne(context.TODO(), exam)
	return result, err
}

func (db *Database) ReadExam(id primitive.ObjectID) (*Exam, error) {
	var exam Exam
	err := db.Exams.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&exam)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}
	return &exam, nil
}

func (db *Database) UpdateExam(id primitive.ObjectID, updatedExam Exam) (*mongo.UpdateResult, error) {
	update := bson.M{
		"$set": updatedExam,
	}

	result, err := db.Exams.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	return result, err
}

func (db *Database) DeleteExam(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	result, err := db.Exams.DeleteOne(context.TODO(), bson.M{"_id": id})
	return result, err
}

func (db *Database) GetAllExams() ([]*Exam, error) {
	findOptions := options.Find()
	var results []*Exam

	cur, err := db.Exams.Find(context.Background(), bson.D{{}}, findOptions)
	if err != nil {
		return results, err
	}

	for cur.Next(context.Background()) {
		var elem Exam
		err := cur.Decode(&elem)
		if err != nil {
			return results, err
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		return results, err
	}

	cur.Close(context.Background())

	return results, nil
}
