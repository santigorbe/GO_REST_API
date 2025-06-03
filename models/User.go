package models

import "time"

// Automaticamente se crean los campos, id, created_at, updated_at, deleted_at

type User struct {
	ID        uint      `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`

	FirstName string `gorm:"not null" json:"first_name" validate:"required,min=2,max=100"`
	LastName  string `gorm:"not null" json:"last_name" validate:"required,min=2,max=100"`
	Email     string `gorm:"not null; unique_index" json:"email" validate:"required,email"`
	Tasks     []Task `json:"tasks"`
}