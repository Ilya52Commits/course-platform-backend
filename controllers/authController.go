package controllers

import (
	"github.com/Ilya52Commits/course-platform/database"
	"github.com/Ilya52Commits/course-platform/models"
	"github.com/Ilya52Commits/course-platform/scripts"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

const SecretKey = "secret" // Секретный ключ для токена пользователя

// MailConfirm /* Реализация функции подтверждения почты */
func MailConfirm(c *fiber.Ctx) error {
	var data map[string]string // Создание хэш-таблицы для данных

	// Проверка тела запроса на ошибку
	if err := c.BodyParser(&data); err != nil {
		// Если ошибка не пуста то происходит выход из функции
		return err
	}

	var user models.User // Возврать модели User

	// Поиск пользователя по почте
	database.DB.Where("is_email_verified = ?", false).First(&user)

	// Если флаг повторной отправки письма не пустой
	if data["repeat"] != "" {
		validationCode := rand.Intn(900000) + 100000 // Генерация случайного числа от 100000 до 999999

		// Вызов метода для отправки письма на почту пользователя
		if err := scripts.SendMail(user.Email, validationCode); err != nil {
			// Если почта некоректная, то происходит выход из функции
			return c.JSON(fiber.Map{
				"message": "Почта не коректная",
			})
		}

		user.VerifiedCode = validationCode // Присваивание нового кода

		// Сохранение изменений
		database.DB.Save(&user)

		// Отправка сообщения о новом письме
		return c.JSON(fiber.Map{
			"message": "Был отправлен повторный код",
		})
	}

	// Если пользователь не был найден
	if strconv.Itoa(user.VerifiedCode) != data["code"] {

		user.VerifiedCode = 0 // Обнуление кода в бд

		// Сохранение изменений
		database.DB.Save(&user)

		// Статус ответа - код 404
		c.Status(fiber.StatusNotFound)
		// Возврат json хэш-таблицы с ошибкой
		return c.JSON(fiber.Map{
			"message": "Код не совпадает",
		})
	}

	user.VerifiedCode = 0       // Обнуление кода
	user.IsEmailVerified = true // Изменение статуса почты на подтверждённый

	// Созранение изменений на почту
	database.DB.Save(&user)

	// Отправка сообщения об успешном подтверждении почты
	return c.JSON(fiber.Map{
		"message": "Почта подтверждена",
	})
}

// Register /* Реализация функции регистрации */
func Register(c *fiber.Ctx) error {
	var data map[string]string // Создание хэш-таблицы для данных

	// Проверка тела запроса на ошибку
	if err := c.BodyParser(&data); err != nil {
		// Если ошибка не пуста, то происходит выход из функции
		return err
	}

	/* Проверка, была ли почта проеверена */
	var user models.User // Возврать модели User
	// Поиск первый результат почты в базе данных
	database.DB.Where("email = ?", data["email"]).First(&user)

	// Если пользователь не был найден
	if user.Id == 0 {
		validationCode := rand.Intn(900000) + 100000 // Генерация случайного числа от 100000 до 999999

		// Вызов метода для отправки письма на почту пользователя
		if err := scripts.SendMail(data["email"], validationCode); err != nil {
			// Если почта некоректная, то происходит выход из функции
			return c.JSON(fiber.Map{
				"message": "Почта не коректная",
			})
		}

		// Хэширование пароля
		password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

		// Создание объекта модели User
		user := models.User{
			Name:            data["name"],   // Применение имени
			Email:           data["email"],  // Применение почты
			Password:        password,       // Применение хэшированного пароля
			VerifiedCode:    validationCode, // Применение кода валидации
			IsEmailVerified: false,
		}

		// Сохраниение объекта в базу данных
		database.DB.Create(&user)

		// Сохранение изменений
		database.DB.Save(&user)

		// Возвращение файла json объекта user
		return c.JSON(user)
	}

	validationCode := rand.Intn(900000) + 100000 // Генерация случайного числа от 100000 до 999999

	// Вызов метода для отправки письма на почту пользователя
	if err := scripts.SendMail(data["email"], validationCode); err != nil {
		// Если почта некоректная, то происходит выход из функции
		return c.JSON(fiber.Map{
			"message": "Почта не коректная",
		})
	}

	user.VerifiedCode = validationCode // Присваивание нового кода

	// Обновление бд
	database.DB.Save(&user)

	// Отправка сообщения об отсутвии подтверждения
	return c.JSON(fiber.Map{
		"message": "Вы не подтвердили почту",
	})
}

// Login /* Реализация функции входа */
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
			"message": "Пользователь не найден",
		})
	}

	if user.IsEmailVerified == false {
		// Статус ответа - код 404
		c.Status(fiber.StatusNotFound)
		// Возврат json хэш-таблицы с ошибкой
		return c.JSON(fiber.Map{
			"message": "Вы не подтвердили почту",
		})
	}

	// Сравнивает хэш пароля user в бд с введёным от пользователя
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		// Статус ответа - код 404
		c.Status(fiber.StatusBadRequest)
		// Возврат json хэш-таблицы с ошибкой
		return c.JSON(fiber.Map{
			"message": "Неправильный пароль",
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
			"message": "Не удалось войти в систему",
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
		"message": "Успешно",
	})
}

// User /* Создание функции получения пользователя */
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
			"message": "неаутентифицированный",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims) // Извлечение утверждений из токена

	var user models.User // Возврать модели User

	// Поиск первого пользователя по id
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	// Возврат json файла найденного пользователя
	return c.JSON(user)
}

// Logout /* Создание функции для выхода из системы */
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
		"message": "успешно",
	})
}
