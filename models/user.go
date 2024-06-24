package models

import "time"

type User struct {
	ID        uint       `gorm:"primary key;autoIncrement" json:"id"`
	Password  *string    `gorm:"not null" json:"password" binding:"required"`
	Email     *string    `gorm:"unique" json:"email" binding:"required"`
	Role      *string    `gorm:"not null" json:"role" binding:"required"`
	Name      *string    `gorm:"not null" json:"name" binding:"required"`
	Info      []Info     `gorm:"foreignKey:UserID"` //FK to info table, one to many relationship
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
