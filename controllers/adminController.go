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
		AuthorId:    user.Id,
	}

	database.DB.Create(&course)

	database.DB.Save(&course)

	return c.JSON(fiber.Map{
		"message": "Курс успешно создан",
	})
}

func CreateModule(c *fiber.Ctx) error {
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

	var course models.Course

	database.DB.Where("name = ?", data["nameCourse"]).First(&course)

	// Поиск первого пользователя по id
	if course.Id == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "курс не найден",
		})
	}

	module := models.Module{
		Name:     data["nameModule"],
		IdCourse: course.Id,
	}

	database.DB.Create(&module)

	database.DB.Save(&module)

	return c.JSON(module)
}

func GetModules(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var course models.Course
	database.DB.Where("name = ?", data["nameCourse"]).First(&course)

	if course.Id == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "курс не найден",
		})
	}

	var modules []models.Module
	database.DB.Where("id_course = ?", course.Id).Find(&modules)

	return c.JSON(modules)
}

func GetCourses(c *fiber.Ctx) error {
	var courses []models.Course
	database.DB.Find(&courses)

	return c.JSON(courses)
}

func CreateLesson(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var module models.Module
	database.DB.Where("name = ?", data["nameModule"]).Find(&module)

	// Поиск первого пользователя по id
	if module.Id == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "модуль не найден",
		})
	}

	lesson := models.Lesson{
		Name:     data["nameLesson"],
		Video:    data["urlVideoLesson"],
		Task:     data["taskLesson"],
		IdModule: module.Id,
	}

	database.DB.Create(&lesson)

	database.DB.Save(&lesson)

	return c.JSON(lesson)
}

func GetLessons(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var module models.Module
	database.DB.Where("name = ?", data["nameModule"]).First(&module)

	if module.Id == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "модуль не найден",
		})
	}

	var lessons []models.Lesson
	database.DB.Where("id_module = ?", module.Id).Find(&lessons)

	return c.JSON(lessons)
}

func GetLesson(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var module models.Module
	database.DB.Where("name = ?", data["nameModule"]).First(&module)

	if module.Id == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "модуль не найден",
		})
	}

	var lesson models.Lesson
	database.DB.Where("id_module = ?", module.Id).First(&lesson)

	return c.JSON(lesson)
}
