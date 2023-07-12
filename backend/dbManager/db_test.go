package dbManager_test

import (
	"context"
	"mockExamSchedulerBackend/dbManager"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func setupDB(t *testing.T) *dbManager.Database {
	err := godotenv.Load("../.env")
	connectionString := dbManager.GetConnectionStringFromEnvFile(".env")
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	db, err := dbManager.NewDatabase(connectionString, "test")
	if err != nil {
		t.Fatal(err)
	}
	// Clear the exams collection for a clean test run.
	_, err = db.Exams.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func TestCreateExam(t *testing.T) {
	db := setupDB(t)
	testId, _ := primitive.ObjectIDFromHex("60e9d0f7a5f9a7b4a0f3b1a1")

	exam := dbManager.Exam{Id: testId, Name: "Math Exam", IsInProgress: true, Challenger: "John", StartTime: 1625626031, EndTime: 1625629631}

	_, err := db.CreateExam(exam)
	assert.NoError(t, err)

	retrievedExam, err := db.ReadExam(testId)
	assert.NoError(t, err)
	assert.Equal(t, exam, *retrievedExam)
}

func TestUpdateExam(t *testing.T) {
	db := setupDB(t)
	testId, _ := primitive.ObjectIDFromHex("60e9d0f7a5f9a7b4a0f3b1a1")

	exam := dbManager.Exam{Id: testId, Name: "Math Exam", IsInProgress: true, Challenger: "John", StartTime: 1625626031, EndTime: 1625629631}
	db.CreateExam(exam)

	exam.Name = "Science Exam"
	_, err := db.UpdateExam(testId, exam)
	assert.NoError(t, err)

	updatedExam, err := db.ReadExam(testId)
	assert.NoError(t, err)
	assert.Equal(t, "Science Exam", updatedExam.Name)
}

func TestGetAllExams(t *testing.T) {
	db := setupDB(t)
	testId, _ := primitive.ObjectIDFromHex("60e9d0f7a5f9a7b4a0f3b1a1")
	testId2, _ := primitive.ObjectIDFromHex("60e9d0f7a5f9a7b4a0f3b1a2")

	exam1 := dbManager.Exam{Id: testId, Name: "Math Exam", IsInProgress: true, Challenger: "John", StartTime: 1625626031, EndTime: 1625629631}
	db.CreateExam(exam1)

	exam2 := dbManager.Exam{Id: testId2, Name: "Science Exam", IsInProgress: false, Challenger: "Alex", StartTime: 1625626041, EndTime: 1625629641}
	db.CreateExam(exam2)

	exams, err := db.GetAllExams()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(exams))
}

func TestDeleteExam(t *testing.T) {
	db := setupDB(t)
	testId, _ := primitive.ObjectIDFromHex("60e9d0f7a5f9a7b4a0f3b1a1")

	exam := dbManager.Exam{Id: testId, Name: "Math Exam", IsInProgress: true, Challenger: "John", StartTime: 1625626031, EndTime: 1625629631}
	db.CreateExam(exam)

	_, err := db.DeleteExam(testId)
	assert.NoError(t, err)

	deletedExam, err := db.ReadExam(testId)
	assert.ErrorIs(t, err, mongo.ErrNoDocuments)
	assert.Nil(t, deletedExam)
}
