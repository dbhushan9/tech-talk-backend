package services

import (
	"dbhushan9/tech-talk-backend/domain"
	"dbhushan9/tech-talk-backend/models"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type TechTalkService struct {
	logger *log.Logger
	dao    *domain.TechTalkDAO
}

func NewTechTalkService(l *log.Logger, db *mongo.Database) *TechTalkService {
	t := new(TechTalkService)
	t.initialize(l, db)
	return t
}

func (t *TechTalkService) initialize(l *log.Logger, db *mongo.Database) {
	t.logger = l
	t.dao = domain.NewTechTalkDAO(l, db)
}

func (t *TechTalkService) GetAll() ([]*models.TechTalk, error) {
	return t.dao.GetAll()
}

func (t *TechTalkService) Create(tt *models.TechTalk) (*models.TechTalk, error) {
	return t.dao.Save(tt)
}

func (t *TechTalkService) Update(id string, tt *models.TechTalk) (*models.TechTalk, error) {
	return t.dao.Update(id, tt)
}
