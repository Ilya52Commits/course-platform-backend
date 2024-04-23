package models

// Модель таблицы пользователя
type User struct {
	Id       int    `json:"id"`    // Идентификатор
	Name     string `json:"name"`  // Имя
	Email    string `json:"email"` //gorm:"unique"` // Почта
	Password []byte `json:"-"`     // Пароль
}
