package Services

import (
	"aviasales"
	"aviasales/Databases"
)

// AnketaService - сервис для работы с анкетами
type AnketaService struct {
	db *Databases.Database
}

// NewAnketaService - конструктор
func NewAnketaService(db *Databases.Database) *AnketaService {
	return &AnketaService{db: db}
}

// CreateAnketa - создание анкеты
func (s *AnketaService) CreateAnketa(anketa main.Anketa) error {
	return s.db.SaveAnketa(anketa)
}

// DeleteAnketa - удаление анкеты по ID
func (s *AnketaService) DeleteAnketa(id int) error {
	return s.db.DeleteAnketa(id)
}

// GetAnketa - получение анкеты по ID
func (s *AnketaService) GetAnketa(id int) (main.Anketa, error) {
	return s.db.GetAnketa(id)
}
