package models

import "time"

type Token struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID    string    `gorm:"not null;type:uuid"`
	Token     string    `gorm:"uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	IsRevoked bool      `gorm:"default:false"`
	TokenType string    `gorm:"not null"`
}
