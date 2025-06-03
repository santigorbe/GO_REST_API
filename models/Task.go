package models

import "time"

type Task struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`

	Title       string `gorm:"type:varchar(100);not null; unique_index" json:"title" validate:"required,min=2,max=100"`
	Description string `json:"description" gorm:"type:text;not null" validate:"required,min=2,max=500"`
	Done        bool   `gorm:"default:false" json:"done"`
	UserID      uint   `json:"user_id"`
}
