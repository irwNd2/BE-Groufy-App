package models

type Info struct {
	ID      uint    `gorm:"primary key;autoIncrement" json:"id"`
	Author  *string `json:"author"`
	Title   *string `json:"title"`
	Content *string `json:"content"`
	Date    *string `json:"date"`
}
