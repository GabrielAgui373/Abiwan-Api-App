package repositories

import (
	"github.com/gabrielagui373/obiwanapp-api/internal/models"
	"gorm.io/gorm"
)

type TopicRepository struct {
	*BaseRepository[models.Topic]
}

func NewTopicRepository(db *gorm.DB) *TopicRepository {
	return &TopicRepository{
		BaseRepository: NewBaseRepository[models.Topic](db),
	}
}

func (r *TopicRepository) findTopics(userID string, additionalConditions ...func(*gorm.DB) *gorm.DB) ([]models.Topic, error) {
	var topics []models.Topic

	query := r.db.
		Preload("Children", r.recursivePreload(userID)).
		Where("user_id = ?", userID).
		Order("topic_order ASC")

	for _, condition := range additionalConditions {
		query = condition(query)
	}

	err := query.Find(&topics).Error
	return topics, err
}

func (r *TopicRepository) FindAll(userID string) ([]models.Topic, error) {
	return r.findTopics(userID, func(db *gorm.DB) *gorm.DB {
		return db.Where("parent_id IS NULL")
	})
}

func (r *TopicRepository) FindByID(userID string, topicID string) (*models.Topic, error) {
	topics, err := r.findTopics(userID, func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", topicID)
	})

	if err != nil {
		return nil, err
	}

	if len(topics) == 0 {
		return nil, nil
	}

	return &topics[0], nil
}

func (r *TopicRepository) recursivePreload(userID string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Where("user_id = ?", userID).
			Order("topic_order ASC").
			Preload("Children", r.recursivePreload(userID))
	}
}
