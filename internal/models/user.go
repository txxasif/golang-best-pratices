package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name" gorm:"not null"`
	Email        string         `json:"email" gorm:"unique;not null"`
	Password     string         `json:"-" gorm:"not null"`
	IsVerified   bool           `json:"is_verified" gorm:"default:false"`
	OTPCode      string         `json:"-" gorm:"size:6"`
	OTPExpiresAt time.Time      `json:"-"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
