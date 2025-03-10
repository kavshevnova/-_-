package Services

import (
	"ankets_and_clients/Databases"
	"ankets_and_clients/Domain"
)

type ClientService struct {
	bd *Databases.Database_clients
}

// конструктор
func NewClientService(bd *Databases.Database_clients) *ClientService {
	return &ClientService{bd: bd}
}

// создание
func (s *ClientService) CreateClient(client Domain.Clients) error {
	return s.bd.SaveClient(client)
}

// удаление
func (s *ClientService) DeleteClient(name string) error {
	return s.bd.DeleteClient(name)
}

// получение
func (s *ClientService) GetClient(name string) (Domain.Clients, error) {
	return s.bd.GetClient(name)
}
