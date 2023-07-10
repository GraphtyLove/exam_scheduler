package mongo

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func setupDB(t *testing.T) *Database {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}
	connection_string := os.Getenv("MONGO_CONNECTION_STRING")
	db, err := New(connection_string, "test")
	if err != nil {
		t.Fatal(err)
	}
	// Clear the exams collection for a clean test run.
	_, err = db.db.Collection("exams").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func TestCreateExam(t *testing.T) {
	db := setupDB(t)

	exam := Exam{Id: 1, Name: "Math Exam", IsInProgress: true, Challenger: "John", StartTime: 1625626031, EndTime: 1625629631}

	_, err := db.CreateExam(exam)
	assert.NoError(t, err)

	retrievedExam, err := db.ReadExam(1)
	assert.NoError(t, err)
	assert.Equal(t, exam, *retrievedExam)
}

func TestUpdateExam(t *testing.T) {
	db := setupDB(t)

	exam := Exam{Id: 1, Name: "Math Exam", IsInProgress: true, Challenger: "John", StartTime: 1625626031, EndTime: 1625629631}
	db.CreateExam(exam)

	exam.Name = "Science Exam"
	_, err := db.UpdateExam(1, exam)
	assert.NoError(t, err)

	updatedExam, err := db.ReadExam(1)
	assert.NoError(t, err)
	assert.Equal(t, "Science Exam", updatedExam.Name)
}

func TestGetAllExams(t *testing.T) {
	db := setupDB(t)

	exam1 := Exam{Id: 1, Name: "Math Exam", IsInProgress: true, Challenger: "John", StartTime: 1625626031, EndTime: 1625629631}
	db.CreateExam(exam1)

	exam2 := Exam{Id: 2, Name: "Science Exam", IsInProgress: false, Challenger: "Alex", StartTime: 1625626041, EndTime: 1625629641}
	db.CreateExam(exam2)

	exams, err := db.GetAllExams()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(exams))
}

func TestDeleteExam(t *testing.T) {
	db := setupDB(t)

	exam := Exam{Id: 1, Name: "Math Exam", IsInProgress: true, Challenger: "John", StartTime: 1625626031, EndTime: 1625629631}
	db.CreateExam(exam)

	_, err := db.DeleteExam(1)
	assert.NoError(t, err)

	deletedExam, err := db.ReadExam(1)
	assert.ErrorIs(t, err, mongo.ErrNoDocuments)
	assert.Nil(t, deletedExam)
}
