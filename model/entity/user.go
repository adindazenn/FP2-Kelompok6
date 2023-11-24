package entity

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey"`
	Username  string    `gorm:"uniqueIndex"`
	Email     string    `gorm:"uniqueIndex"`
	Password  string    
	Age       int    
	CreatedAt time.Time 
	UpdatedAt time.Time 
}
