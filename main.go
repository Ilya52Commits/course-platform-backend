package main

import (
	"github.com/Ilya52Commits/course-platform/database"
	"github.com/Ilya52Commits/course-platform/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Вызов функции Connect в пакете database
	database.Connect()

	app := fiber.New() // Создание экземпляра fiber для работы с HTTP запросами

	// Добавляется middleware для обработки CORS-заголовков с помощью функции
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept",
		// Разрешаются запросы с http://localhost:5173
		AllowOrigins: "http://localhost:5173",
		// Разрешены кросс-доменные запросы с использованием куки
		AllowCredentials: true,
	}))

	// Настраивает маршруты приложения
	routers.SetUp(app)

	// Настраивается порт 8000
	err := app.Listen(":8000")
	if err != nil {
		return
	}
}
