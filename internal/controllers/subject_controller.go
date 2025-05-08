package controllers

import (
	"github.com/gabrielagui373/obiwanapp-api/internal/models"
	"github.com/gabrielagui373/obiwanapp-api/internal/services"
)

type SubjectController struct {
	*BaseController[models.Subject]
	service *services.SubjectService
}

func NewSubjectController(service *services.SubjectService) *SubjectController {
	return &SubjectController{
		BaseController: NewBaseController(service),
		service:        service,
	}
}
