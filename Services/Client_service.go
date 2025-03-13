package Services

import (
	"ankets_and_clients/Databases"
	"ankets_and_clients/Domain"
	logger "ankets_and_clients/internal/package"
)

type ClientService struct {
	bd     *Databases.Database_clients
	logger *logger.Logger
}

// конструктор
func NewClientService(bd *Databases.Database_clients, logger *logger.Logger) *ClientService {
	return &ClientService{bd: bd, logger: logger}
}

// создание
func (s *ClientService) CreateClient(client Domain.Clients) error {
	s.logger.Log("Entering CreateClient")
	return s.bd.SaveClient(client)
}

// удаление
func (s *ClientService) DeleteClient(name string) error {
	s.logger.Log("Entering DeleteClient")
	return s.bd.DeleteClient(name)
}

// получение
func (s *ClientService) GetClient(name string) (Domain.Clients, error) {
	s.logger.Log("Entering GetAnketa")
	return s.bd.GetClient(name)
}
