package Databases

import (
	"aviasales"
	"errors"
	"sync"
)

// Database - хранилище данных для анкет
type Database struct {
	mu     sync.RWMutex
	ankets map[int]main.Anketa
}

func (db *Database) init() {
	anketa := main.Anketa{
		Name:        "Мишель",
		Id:          123,
		City:        "Moscow",
		Age:         25,
		Weight:      49,
		Height:      175,
		Boobs:       2,
		HairColor:   "Блондинка",
		Nationality: "Русская",
		District:    "ЦСКА",
		Price:       20000,
	}
	db.ankets[123] = anketa
}

// NewDatabase - конструктор (для создания и инициализации нового экземпляра Database)
func NewDatabase() *Database {
	db := &Database{
		ankets: make(map[int]main.Anketa),
	}
	db.init()
	return db
}

// Метод SaveAnketa - сохранение анкеты в хранилище (карте ankets)
func (db *Database) SaveAnketa(anketa main.Anketa) error {
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
func (db *Database) GetAnketa(id int) (main.Anketa, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	anketa, exists := db.ankets[id]
	if !exists {
		return main.Anketa{}, ErrAnketaNotFound
	}

	return anketa, nil
}

// Ошибка, если анкета не найдена
var ErrAnketaNotFound = errors.New("анкета не найдена")
