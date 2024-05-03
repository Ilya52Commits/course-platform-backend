package models

// Module /* Модель таблицы модуля */
type Module struct {
	Id          int    `json:"id"`        // Идентификатор
	Name        string `json:"name"`      // Имя
	Lessons     []int  `json:"lessons"`   // Массив id уроков
	Description string `json:"author"`    // Описание
	IdCourse    int    `json:"id_course"` // Идентификатор курса
}
