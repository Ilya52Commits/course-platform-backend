package models

// Course /* Модель таблицы курса */
type Course struct {
	Id          int    `json:"id"`          // Идентификатор
	Name        string `json:"name"`        // Название
	Description string `json:"description"` // Описание
	Author      string `json:"author"`      // Создатель курса
	//IdModules   []int  `json:"id_modules"`  // Массив id модулей
}
