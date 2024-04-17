package routers

import (
	"github.com/Ilya52Commits/course-platform/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/apo/login", controllers.Login)
}
