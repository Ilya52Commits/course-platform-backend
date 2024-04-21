package main

import (
	"fmt"
	"github.com/Ilya52Commits/course-platform/database"
	"github.com/Ilya52Commits/course-platform/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"net/smtp"
)

// Реализация функции подтверждения почты ************************
func SendMail() {
	// Информация об отправителе
	from := "krasnenkov.ilia@gmail.com"
	password := "s5067a301"

	// Информация о получателе
	to := []string{
		"krasnenkov.ilya@inbox.ru",
	}

	// smtp сервер конфигурация
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Сообщение.
	message := []byte("Тестовой сообщение через golang.")

	// Авторизация.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Отправка почты.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Почта отправлена!")
}

func main() {
	// Проверка отправки сообщения **********************
	SendMail() // Гугл запрашивает дополнительное подтверждение

	// Вызов функции Connect в паете database
	database.Connect()

	app := fiber.New() // Создание экземлпяра fiber для работы с HTTP запросами

	// Добавляется middleware для обработки CORS-заголовков с помощью функции
	app.Use(cors.New(cors.Config{
		// Разрешаются запросы с http://localhost:5173
		AllowOrigins: "http://localhost:5173",
		// Разрешены кросс-доменные запросы с использованием куки
		AllowCredentials: true,
	}))

	// Настраивает маршруты приложения
	routers.SetUp(app)

	// Начтраивается порт 8000
	app.Listen(":8000")
}
