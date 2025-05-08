package models

type Subject struct {
	ID          string  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string  `gorm:"type:text;not null" json:"name"`
	Description *string `gorm:"type:text" json:"description,omitempty"`
	Color       string  `gorm:"type:valid_colors" json:"color"`
	UserID      string  `gorm:"not null;index" json:"-"`
	User        User    `gorm:"foreignKey:UserID" json:"-"`
}
