package main

import (
	"log"

	"mockExamSchedulerBackend/dbManager"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	connectionString := dbManager.GetConnectionStringFromEnvFile(".env")
	db, err := dbManager.NewDatabase(connectionString, "production")
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Alive")
	})

	app.Post("/exam", func(c *fiber.Ctx) error {
		exam := new(dbManager.Exam)

		if err := c.BodyParser(exam); err != nil {
			log.Println(err)
			return c.Status(400).SendString(err.Error())
		}

		_, err := db.CreateExam(*exam)
		if err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(exam)
	})

	app.Get("/exam/:id", func(c *fiber.Ctx) error {
		id, err := primitive.ObjectIDFromHex(c.Params("id"))

		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.Status(400).SendString("{'error': 'No exam with that id'}")
			}
			log.Println(err)
			return c.Status(400).SendString(err.Error())
		}

		exam, err := db.ReadExam(id)
		if err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(exam)
	})

	app.Put("/exam/:id", func(c *fiber.Ctx) error {
		id, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			log.Printf("Error while parsing id: %s", err)
			return c.Status(400).SendString(err.Error())
		}

		updatedExam := new(dbManager.Exam)
		if err := c.BodyParser(updatedExam); err != nil {
			log.Println(err)
			return c.Status(400).SendString(err.Error())
		}

		_, err = db.UpdateExam(id, *updatedExam)
		if err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(updatedExam)
	})

	app.Delete("/exam/:id", func(c *fiber.Ctx) error {
		id, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			log.Println(err)
			return c.Status(400).SendString(err.Error())
		}

		_, err = db.DeleteExam(id)
		if err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		return c.SendString("Deleted")
	})

	app.Get("/exams", func(c *fiber.Ctx) error {
		exams, err := db.GetAllExams()
		if err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(exams)
	})

	app.Listen(":3000")
}
