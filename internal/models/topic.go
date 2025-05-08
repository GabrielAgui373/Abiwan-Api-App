package models

type TopicStatus string

const (
	StatusNotSeen  TopicStatus = "not_seen"
	StatusSeen     TopicStatus = "seen"
	StatusReviewed TopicStatus = "reviewed"
)

type Topic struct {
	ID          string      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name        string      `gorm:"type:text;not null" json:"name"`
	Description *string     `gorm:"type:text" json:"description,omitempty"`
	SubjectID   string      `gorm:"not null;index" json:"subject_id"`
	Subject     Subject     `gorm:"foreignKey:SubjectID" json:"-"`
	ParentID    *string     `gorm:"index" json:"parent_id,omitempty"`
	Parent      *Topic      `gorm:"foreignKey:ParentID" json:"-"`
	Children    []Topic     `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	Status      TopicStatus `gorm:"type:text" json:"status"`
	TopicOrder  int         `gorm:"default:0" json:"topic_order"`
	UserID      string      `gorm:"not null;index" json:"-"`
	User        User        `gorm:"foreignKey:UserID" json:"-"`
}
