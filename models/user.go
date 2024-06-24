package models

import "time"

type User struct {
	ID        uint       `gorm:"primary key;autoIncrement" json:"id"`
	Username  *string    `json:"username"`
	Password  *string    `json:"password"`
	Email     *string    `json:"email"`
	Role      *string    `json:"role"`
	Name      *string    `json:"name"`
	Info      []Info       `gorm:"foreignKey:UserID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
