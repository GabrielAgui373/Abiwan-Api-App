package models

type Role struct {
	ID          string       `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string       `gorm:"uniqueIndex;not null"`
	Description string       `gorm:"not null"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}
