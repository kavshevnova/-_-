package main

type ClientService struct {
	bd *Database_clients
}

// конструктор
func NewClientService(bd *Database_clients) *ClientService {
	return &ClientService{bd: bd}
}

// создание
func (s *ClientService) CreateClient(client Clients) error {
	return s.bd.SaveClient(client)
}

// удаление
func (s *ClientService) DeleteClient(name string) error {
	return s.bd.DeleteClient(name)
}

// получение
func (s *ClientService) GetClient(name string) (Clients, error) {
	return s.bd.GetClient(name)
}
