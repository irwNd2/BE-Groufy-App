package models

import "time"

type Info struct {
	ID        uint       `gorm:"primary key;autoIncrement" json:"id"`
	Author    *string    `json:"author"`
	Title     *string    `json:"title"`
	Content   *string    `json:"content"`
	Date      *string    `json:"date"`
	UserID    uint       `json:"user_id"` //key to join with user table
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
