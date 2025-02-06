package models

import "gorm.io/gorm"

// Article represents an article model
type Article struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255);not null"`
	Content string `gorm:"type:text;not null"`
}
