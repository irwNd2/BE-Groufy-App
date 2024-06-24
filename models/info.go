package models

import "time"

type Info struct {
	ID        uint       `gorm:"primary key;autoIncrement" json:"id"`
	Author    *string    `gorm:"not null" json:"author" binding:"required"`
	Title     *string    `gorm:"not null" json:"title" binding:"required"`
	Content   *string    `gorm:"not null" json:"content" binding:"required"`
	UserID    uint       `json:"user_id"` //key to join with user table
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
