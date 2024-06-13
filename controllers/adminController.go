package controllers

import (
	"github.com/Ilya52Commits/course-platform/database"
	"github.com/Ilya52Commits/course-platform/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)


func CreateCourse(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt") // Извлечение куки

	// Проверка токена
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Возврат срез байтов секретного ключа
		return []byte(SecretKey), nil
	})

	// Если при разборе токена вышла ошибка
	if err != nil {
		// Статус ответа - код 404
		c.Status(fiber.StatusUnauthorized)
		// Возврат json хэш-таблицы с ошибкой
		//goland:noinspection LanguageDetectionInspection
		return c.JSON(fiber.Map{
			"message": "не аутентифицированный",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims) // Извлечение утверждений из токена

	var user models.User // Возврат модели User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	// Поиск первого пользователя по id
	if user.Id == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "проблемы с аккаунтом",
		})
	}

	var data map[string]string // Создание хэш-таблицы для данных

	// Проверка тела запроса на ошибку
	if err = c.BodyParser(&data); err != nil {
		// Если ошибка не пуста, то происходит выход из функции
		return err
	}

	course := models.Course{
		Name:        data["nameCourse"],
		Description: data["descriptionCourse"],
		Author:      user.Email,
	}

	database.DB.Create(&course)

	database.DB.Save(&course)

	return c.JSON(fiber.Map{
		"message": "Курс успешно создан",
	})
}

// Токен уже распознан, попробовать передавать почту в app.jsx
// Почитать, не нарушает ли это правило токенизации