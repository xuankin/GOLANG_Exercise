package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primary_key;auto_increment"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
