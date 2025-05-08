package services

import (
	"github.com/gabrielagui373/obiwanapp-api/internal/models"
	"github.com/gabrielagui373/obiwanapp-api/internal/repositories"
)

type SubjectService struct {
	*BaseService[models.Subject]
	repo *repositories.SubjectRepository
}

func NewSubjectService(repo *repositories.SubjectRepository) *SubjectService {
	return &SubjectService{
		BaseService: NewBaseService(repo),
		repo:        repo,
	}
}
