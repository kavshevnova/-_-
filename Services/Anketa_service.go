package Services

import (
	"ankets_and_clients/Databases"
	"ankets_and_clients/Domain"
	logger "ankets_and_clients/internal/package"
)

// AnketaService - сервис для работы с анкетами
type AnketaService struct {
	db     *Databases.Database
	logger *logger.Logger
}

// NewAnketaService - конструктор
func NewAnketaService(db *Databases.Database, logger *logger.Logger) *AnketaService {
	return &AnketaService{db: db, logger: logger}
}

// CreateAnketa - создание анкеты
func (s *AnketaService) CreateAnketa(anketa Domain.Anketa) error {
	s.logger.Log("Entering CreateAnketa")
	return s.db.SaveAnketa(anketa)
}

// DeleteAnketa - удаление анкеты по ID
func (s *AnketaService) DeleteAnketa(id int) error {
	s.logger.Log("Entering DeleteAnketa")
	return s.db.DeleteAnketa(id)
}

// GetAnketa - получение анкеты по ID
func (s *AnketaService) GetAnketa(id int) (Domain.Anketa, error) {
	s.logger.Log("Entering GetAnketa")
	return s.db.GetAnketa(id)
}
