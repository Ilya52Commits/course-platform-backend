package models

// User /* Модель таблицы пользователя */
type User struct {
	Id              int    `json:"id"`                  // Идентификатор
	Name            string `json:"name"`                // Имя
	Email           string `json:"email" gorm:"unique"` // Почта
	Password        []byte `json:"-"`                   // Пароль
	VerifiedCode    int    `json:"verified_code"`
	IsEmailVerified bool   `json:"is_email_verified" gorm:"default:false;not null"` // Флаг подтверждения почты
	IsAdmin         bool   `json:"is_admin" gorm:"default:false;not null"`          // Флаг подтверждения на создателя курса
	//CreatedCourseId []int  `json:"created_course_id"`                               // Массив идентификаторов созданных курсов
	//TrackedCourseId []int  `json:"tracked_course_id"`                               // Массив идентификаторов созданных курсов
}
