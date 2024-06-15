package models

// Course /* Модель таблицы курса */
type Course struct {
	Id          int    	`json:"id"`          	// Идентификатор
	Name        string 	`json:"name"`        	// Название
	Description string 	`json:"description"` 	// Описание
	AuthorId	int		`json:"authorId"`      	// Создатель курса
}
