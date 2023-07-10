package main

import (
	"log"
	"strconv"

	"mockExamSchedulerBackend/mongo"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	connectionString := mongo.GetConnectionStringFromEnvFile(".env")
	db, err := mongo.New(connectionString, "production")
	if err != nil {
		log.Fatal(err)
	}

	app.Post("/exam", func(c *fiber.Ctx) error {
		exam := new(mongo.Exam)

		if err := c.BodyParser(exam); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		_, err := db.CreateExam(*exam)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(exam)
	})

	app.Get("/exam/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}

		exam, err := db.ReadExam(id)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(exam)
	})

	app.Put("/exam/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}

		updatedExam := new(mongo.Exam)
		if err := c.BodyParser(updatedExam); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		_, err = db.UpdateExam(id, *updatedExam)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(updatedExam)
	})

	app.Delete("/exam/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}

		_, err = db.DeleteExam(id)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.SendString("Deleted")
	})

	app.Get("/exams", func(c *fiber.Ctx) error {
		exams, err := db.GetAllExams()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(exams)
	})

	app.Listen(":3000")
}
