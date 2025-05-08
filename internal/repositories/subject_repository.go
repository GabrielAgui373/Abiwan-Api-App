package repositories

import (
	"github.com/gabrielagui373/obiwanapp-api/internal/models"
	"gorm.io/gorm"
)

type SubjectRepository struct {
	*BaseRepository[models.Subject]
}

func NewSubjectRepository(db *gorm.DB) *SubjectRepository {
	return &SubjectRepository{
		BaseRepository: NewBaseRepository[models.Subject](db),
	}
}
