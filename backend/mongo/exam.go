package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Exam struct {
	Id           int    `bson:"_id,omitempty"`
	Name         string `bson:"name"`
	IsInProgress bool   `bson:"is_in_progress"`
	Challenger   string `bson:"challenger"`
	StartTime    int    `bson:"start_time"`
	EndTime      int    `bson:"end_time"`
}

func (db *Database) CreateExam(exam Exam) (*mongo.InsertOneResult, error) {
	result, err := db.Exams.InsertOne(context.TODO(), exam)
	return result, err
}

func (db *Database) ReadExam(id int) (*Exam, error) {
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

func (db *Database) UpdateExam(id int, updatedExam Exam) (*mongo.UpdateResult, error) {
	update := bson.M{
		"$set": updatedExam,
	}

	result, err := db.Exams.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	return result, err
}

func (db *Database) DeleteExam(id int) (*mongo.DeleteResult, error) {
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
