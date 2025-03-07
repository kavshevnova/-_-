package main

// AnketaService - сервис для работы с анкетами
type AnketaService struct {
	db *Database
}

// NewAnketaService - конструктор
func NewAnketaService(db *Database) *AnketaService {
	return &AnketaService{db: db}
}

// CreateAnketa - создание анкеты
func (s *AnketaService) CreateAnketa(anketa Anketa) error {
	return s.db.SaveAnketa(anketa)
}

// DeleteAnketa - удаление анкеты по ID
func (s *AnketaService) DeleteAnketa(id int) error {
	return s.db.DeleteAnketa(id)
}

// GetAnketa - получение анкеты по ID
func (s *AnketaService) GetAnketa(id int) (Anketa, error) {
	return s.db.GetAnketa(id)
}
