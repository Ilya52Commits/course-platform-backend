package models

// User /* Модель таблицы пользователя */
type User struct {
	Id              int    `json:"id"`                  // Идентификатор
	Name            string `json:"name"`                // Имя
	Email           string `json:"email" gorm:"unique"` // Почта
	Password        []byte `json:"-"`                   // Пароль
	VerifiedCode    int    `json:"verified_code"`		// Код подтверждения почты
	IsEmailVerified bool   `json:"is_email_verified" gorm:"default:false;not null"` // Флаг подтверждения почты
}
