package Controllers

import (
	"ankets_and_clients/Domain"
	"ankets_and_clients/Services"
	logger "ankets_and_clients/internal/package"
	"encoding/json"
	"net/http"
	"strconv"
)

type AnketaController struct {
	service *Services.AnketaService
	logger *logger.Logger
}

func NewAnketaController(service *Services.AnketaService, logger *logger.Logger) *AnketaController {
	return &AnketaController{service: service, logger: logger}
}

// CreateAnketaHandler - обработчик для создания анкеты
func (c *AnketaController) CreateAnketaHandler(w http.ResponseWriter, r *http.Request) {
	c.logger.Log("Entering CreateAnketaHandler")
	var anketa Domain.Anketa
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
	c.logger.Log("Entering DeleteAnketaHandler")
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
	c.logger.Log("Entering GetAnketaHandler")
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

	go install github.com/swaggo/swag/cmd/swag@latest	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(anketa)
}
