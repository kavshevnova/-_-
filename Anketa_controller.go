package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// AnketaController - контроллер для работы с анкетами
type AnketaController struct {
	service *AnketaService
}

// NewAnketaController - конструктор
func NewAnketaController(service *AnketaService) *AnketaController {
	return &AnketaController{service: service}
}

// CreateAnketaHandler - обработчик для создания анкеты
func (c *AnketaController) CreateAnketaHandler(w http.ResponseWriter, r *http.Request) {
	var anketa Anketa
	if err := json.NewDecoder(r.Body).Decode(&anketa); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.service.CreateAnketa(anketa); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Анкета создана"))
}

// DeleteAnketaHandler - обработчик для удаления анкеты
func (c *AnketaController) DeleteAnketaHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := c.service.DeleteAnketa(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Анкета удалена"))
}

// GetAnketaHandler - обработчик для получения анкеты
func (c *AnketaController) GetAnketaHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Не нвйден ID", http.StatusBadRequest)
		return
	}

	anketa, err := c.service.GetAnketa(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(anketa)
}
