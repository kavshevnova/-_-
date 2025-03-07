package main

import (
	"errors"
	"sync"
)

// Database - хранилище данных для анкет
type Database struct {
	mu     sync.RWMutex
	ankets map[int]Anketa
}

// NewDatabase - конструктор (для создания и инициализации нового экземпляра Database)
func NewDatabase() *Database {
	return &Database{
		ankets: make(map[int]Anketa),
	}
}

// Метод SaveAnketa - сохранение анкеты в хранилище (карте ankets)
func (db *Database) SaveAnketa(anketa Anketa) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.ankets[anketa.Id] = anketa
	return nil
}

// DeleteAnketa - удаление анкеты по ID
func (db *Database) DeleteAnketa(id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.ankets[id]; !exists {
		return ErrAnketaNotFound
	}
	delete(db.ankets, id)
	return nil
}

// GetAnketa - получение анкеты по ID
func (db *Database) GetAnketa(id int) (Anketa, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	anketa, exists := db.ankets[id]
	if !exists {
		return Anketa{}, ErrAnketaNotFound
	}

	return anketa, nil
}

// Ошибка, если анкета не найдена
var ErrAnketaNotFound = errors.New("анкета не найдена")
