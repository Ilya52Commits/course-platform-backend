package models

// Lesson /* Модель таблицы урока */
type Lesson struct {
	Id       int    `json:"id"`                                 // Идентификатор
	Name     string `json:"name"`                               // Имя
	Video    string `json:"video"`                              // Название видео
	Task     string `json:"task"`                               // Задание
	IdModule int    `json:"id_module" gorm:"foreignKey:Module"` // Идентификатор курса
}
