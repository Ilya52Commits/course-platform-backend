package controllers

import (
	"github.com/Ilya52Commits/course-platform/database"
	"github.com/Ilya52Commits/course-platform/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

const SecretKey = "secret" // Секретный ключ для токена пользователя

// Реализация функции регистрации
func Register(c *fiber.Ctx) error {
	var data map[string]string // Создание хэш-таблицы для данных

	// Проверка тела запроса на ошибку
	if err := c.BodyParser(&data); err != nil {
		// Если ошибка не пуста то происходит выход из функции
		return err
	}

	// Хэширование пароля
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	// Создание объекта модели User
	user := models.User{
		Name:     data["name"],  // Применение имени
		Email:    data["email"], // Применение почты
		Password: password,      // Применение хэшированного пароля
	}

	// Сохраниение объекта в базу данных
	database.DB.Create(&user)

	// Возвращение файла json объекта user
	return c.JSON(user)
}

// Реализация функции входа
func Login(c *fiber.Ctx) error {
	var data map[string]string // Создание хэш-таблицы для данных

	// Проверка тела запроса на ошибку
	if err := c.BodyParser(&data); err != nil {
		// Если ошибка не пуста то происходит выход из функции
		return err
	}

	var user models.User // Возврать модели User

	// Поиск первый результат почты в базе данных
	database.DB.Where("email = ?", data["email"]).First(&user)

	// Если пользователь не был найден
	if user.Id == 0 {
		// Статус ответа - код 404
		c.Status(fiber.StatusNotFound)
		// Возврат json хэш-таблицы с ошибкой
		return c.JSON(fiber.Map{
			"сообщение": "пользователь не найден",
		})
	}

	// Сравнивает хэш пароля user в бд с введёным от пользователя
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		// Статус ответа - код 404
		c.Status(fiber.StatusBadRequest)
		// Возврат json хэш-таблицы с ошибкой
		return c.JSON(fiber.Map{
			"сообщение": "неправильный пароль",
		})
	}

	// Создание объекта токена
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(user.Id),                 // Страка (эмитет) id преобразованный в строку
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Время истечения токена (1 день)
	})

	token, err := claims.SignedString([]byte(SecretKey)) // Создание токена

	// Если токен не удалось создать
	if err != nil {
		// Статус ответа - код 404
		c.Status(fiber.StatusInternalServerError)
		// Возврат json хэш-таблицы с ошибкой
		return c.JSON(fiber.Map{
			"сообщение": "не удалось войти в систему",
		})
	}

	// Создание куки-файла
	cookie := fiber.Cookie{
		Name:     "jwt",                          // Имя куки
		Value:    token,                          // Содержание куки
		Expires:  time.Now().Add(time.Hour * 24), // Дата истечения куки (1 день)
		HTTPOnly: true,                           // Куки будет доступна только для HTTP
	}

	// Присваивание файла-куки
	c.Cookie(&cookie)

	// Возврать сообщение об успешном входе
	return c.JSON(fiber.Map{
		"сообщение": "успешно",
	})
}

// Создание функции получения пользователя
func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt") // Извлечение куки

	// Проверка токена
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Возврат срез байтов секретного ключа
		return []byte(SecretKey), nil
	})

	// Если при разбора токена вышла ошибка
	if err != nil {
		// Статус ответа - код 404
		c.Status(fiber.StatusUnauthorized)
		// Возврат json хэш-таблицы с ошибкой
		return c.JSON(fiber.Map{
			"сообщение": "неаутентифицированный",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims) // Извлечение утверждений из токена

	var user models.User // Возврать модели User

	// Поиск первого пользователя по id
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	// Возврат json файла найденного пользователя
	return c.JSON(user)
}

// Создание функции для выхода из системы
func Logout(c *fiber.Ctx) error {
	// Обнуление куки-файла
	cookie := fiber.Cookie{
		Name:     "jwt",                      // Имя куки
		Value:    "",                         // Обнуление содержимого
		Expires:  time.Now().Add(-time.Hour), // Обнуление времени куки
		HTTPOnly: true,                       // Куки будет доступна только для HTTP
	}

	// Присваивание куки
	c.Cookie(&cookie)

	// Возврат сообщения об успешном выходе
	return c.JSON(fiber.Map{
		"сообщение": "успешно",
	})
}
