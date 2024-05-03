package models

// Lesson /* Модель таблицы урока */
type Lesson struct {
	Id          int    `json:"id"`                                // Идентификатор
	Name        string `json:"name"`                              // Имя
	Video       string `json:"video"`                             // Название видео
	Description string `json:"author"`                            // Описание
	IdModule    int    `json:"idCourse" gorm:"foreignKey:Module"` // Идентификатор курса
}
