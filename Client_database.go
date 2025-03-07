package main

import (
	"errors"
	"sync"
)

// Database - хранилище данных для клиентов
type Database_clients struct {
	mu      sync.RWMutex
	clients map[string]Clients
}

// Конструктор (для создания и инициализации нового экземпляра Database)
func NewDatabase_clients() *Database_clients {
	return &Database_clients{clients: make(map[string]Clients)}
}

// сохранение пользователя в хранилище (карте clients)
func (bd *Database_clients) SaveClient(client Clients) error {
	bd.mu.Lock()
	defer bd.mu.Unlock()

	bd.clients[client.Name] = client
	return nil
}

// DeleteClient - удаление пользователя по имени
func (bd *Database_clients) DeleteClient(name string) error {
	bd.mu.Lock()
	defer bd.mu.Unlock()

	if _, exists := bd.clients[name]; !exists {
		return ErrClientNotFound
	}
	delete(bd.clients, name)
	return nil
}

// GetAnketa - получение пользователя по имени
func (bd *Database_clients) GetClient(name string) (Clients, error) {
	bd.mu.RLock()
	defer bd.mu.RUnlock()

	client, exists := bd.clients[name]
	if !exists {
		return Clients{}, ErrClientNotFound
	}

	return client, nil
}

// Ошибка, если анкета не найдена
var ErrClientNotFound = errors.New("Рюкзачник не найден")
