package models

// Module /* Модель таблицы модуля */
type Module struct {
	Id          int    `json:"id"`        // Идентификатор
	Name        string `json:"name"`      // Имя
	IdCourse    int    `json:"id_course"` // Идентификатор курса
}
