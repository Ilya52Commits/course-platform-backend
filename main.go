package main

import (
	"github.com/Ilya52Commits/course-platform/database"
	"github.com/Ilya52Commits/course-platform/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	var app = fiber.New()

	routers.SetUp(app)

	app.Listen(":8000")
}
