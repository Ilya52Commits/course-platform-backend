package models

// Course /* Модель таблицы курса */
type Course struct {
	Id          int    `json:"id"`          // Идентификатов
	Name        string `json:"name"`        // Название
	Description string `json:"description"` // Описание
	IdModules   []int  `json:"id_modules"`  // Массив id модулей
	Author      int    `json:"author"`      // Создатель курса
}
