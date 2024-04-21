package main

import (
	"encoding/base64"
	"github.com/Ilya52Commits/course-platform/database"
	"github.com/Ilya52Commits/course-platform/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"net/smtp"
	"strings"
)

// Реализация функции подтверждения почты
//func checkMail() error {
//	auth := smtp.PlainAuth("", "krasnenkov.ilia@gmail.com", "s5067a301", "smtp.gmail.com")
//
//	to := []string{"krasnenkov.ilya@inbox.ru"}
//	msg := []byte("To: " + "krasnenkov.ilya@inbox.ru" + "\r\n" +
//		"Subject: discount Gophers!\r\n" +
//		"\r\n" +
//		"This is the email body.\r\n")
//	err := smtp.SendMail("smtp.gmail.com:465", auth, "krasnenkov.ilia@gmail.com", to, msg)
//	if err != nil {
//		log.Fatal(err)
//		return err
//	}
//	return nil
//}

func SendMail(addr, from, subject, body string, to []string) error {
	r := strings.NewReplacer("\r\n", "", "\r", "", "\n", "", "%0a", "", "%0d", "")

	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	if err = c.Mail(r.Replace(from)); err != nil {
		return err
	}
	for i := range to {
		to[i] = r.Replace(to[i])
		if err = c.Rcpt(to[i]); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	msg := "To: " + strings.Join(to, ",") + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"Content-Transfer-Encoding: base64\r\n" +
		"\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	_, err = w.Write([]byte(msg))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

func main() {
	//err := checkMail()
	//if err != nil {
	//	log.Fatal(err)
	//}
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

	err := SendMail("localhost:8000", "krasnenkov.ilia@gmail.com",
		"Subject text", "Body text", []string{"krasnenkov.ilya@inbox.ru"})
	if err != nil {
		log.Fatal(err)
	}
}
