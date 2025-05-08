package models

type User struct {
	ID            string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Email         string `gorm:"uniqueIndex;not null"`
	PasswordHash  string `gorm:"not null"`
	IsActive      bool   `gorm:"default:true"`
	EmailVerified bool   `gorm:"default:true"`
	Roles         []Role `gorm:"many2many:user_roles;"`
}
