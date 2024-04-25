package routers

import (
	"github.com/Ilya52Commits/course-platform/controllers"
	"github.com/gofiber/fiber/v2"
)

// Создание метода для роутинга
func SetUp(app *fiber.App) {
	// Вызов запроса Post для регистрации
	app.Post("/api/register", controllers.Register)

	// Вызов запроса Post для подтверждения почты
	app.Post("/api/mail-confirm", controllers.MailConfirm)

	// Вызов запроса Post для входа в систему
	app.Post("/api/login", controllers.Login)

	// Вызов запроса Get для вывода пользователя
	app.Get("/api/user", controllers.User)

	// Вызов запроса Post для выхода из системы
	app.Post("/api/logout", controllers.Logout)
}
