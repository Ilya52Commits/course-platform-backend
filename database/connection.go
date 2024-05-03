package database

import (
	"github.com/Ilya52Commits/course-platform/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB Создание экземпляра gorm
var DB *gorm.DB

// Connect /* Создание функции подключения */
func Connect() {
	// Определение dsn (Data Source Name)* - содержит информацию о типе источника данных,
	// его расположении и других параметрах, необходимых для установки соединения
	dsn := "host=localhost user=postgres password=52 dbname=course_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	// Открытие соединения к базе данных по dsn с помощью метода Open
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Если ошибка не равна нулю
	if err != nil {
		// Сообщается об ошибке
		panic("[ERROR]: could not connect to the database | не удалось подключиться к базе данных")
	}

	DB = connection // Присваивание результатов подключения базы данных в экземпляр DB

	// Миграция модели
	err = connection.AutoMigrate(&models.User{},
		&models.Course{},
		&models.Module{},
		&models.Lesson{})
	if err != nil {
		return
	}
}
