package Controllers

import (
	"ankets_and_clients/Domain"
	"ankets_and_clients/Services"
	logger "ankets_and_clients/internal/package"
	"encoding/json"
	"net/http"
)

// контроллер для работы с пользователями
type ClientController struct {
	service *Services.ClientService
	logger  *logger.Logger
}

// Конструктор
func NewClientController(service *Services.ClientService, logger *logger.Logger) *ClientController {
	return &ClientController{service: service, logger: logger}
}

// обработчик для создания
func (c *ClientController) CreateClientHandler(w http.ResponseWriter, r *http.Request) {
	c.logger.Log("Entering CreateClientHandler")
	var client Domain.Clients
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.service.CreateClient(client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Рюкзачник создан"))
}

// обработчик для удаления
func (c *ClientController) DeleteClientHandler(w http.ResponseWriter, r *http.Request) {
	c.logger.Log("Entering DeleteClientHandler")
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Рюкзачник не найден", http.StatusBadRequest)
		return
	}

	if err := c.service.DeleteClient(name); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Рюкзачник удален"))
}

// GetAnketaHandler - обработчик для получения анкеты
func (c *ClientController) GetClientHandler(w http.ResponseWriter, r *http.Request) {
	c.logger.Log("Entering GetClientHandler")
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Рюкзачник не найден", http.StatusBadRequest)
		return
	}

	client, err := c.service.GetClient(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client)
}
