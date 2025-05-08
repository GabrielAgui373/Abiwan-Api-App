package services

import (
	"github.com/gabrielagui373/obiwanapp-api/internal/models"
	"github.com/gabrielagui373/obiwanapp-api/internal/repositories"
	"github.com/gin-gonic/gin"
)

type TopicService struct {
	*BaseService[models.Topic]
	repo *repositories.TopicRepository
}

func NewTopicService(repo *repositories.TopicRepository) *TopicService {
	return &TopicService{
		BaseService: NewBaseService(repo),
		repo:        repo,
	}
}

func (ts *TopicService) GetAll(c *gin.Context) ([]models.Topic, error) {
	user, err := ts.getUser(c)
	if err != nil {
		return nil, err
	}
	return ts.repo.FindAll(user.ID)
}

func (ts *TopicService) GetItemByID(id string, c *gin.Context) (*models.Topic, error) {
	user, err := ts.getUser(c)
	if err != nil {
		return nil, err
	}
	return ts.repo.FindByID(user.ID, id)
}
