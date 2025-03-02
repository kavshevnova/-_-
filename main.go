package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// создать структуру на 6 полей и в одном из хендлеров отдавать структуру
type marshrut struct {
	price int `json:"price"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Получаем значение параметра "question" из запроса
	question := r.URL.Query().Get("question")

	// Проверяем, передан ли параметр "question"
	if question == "" {
		http.Error(w, "Параметр 'question' отсутствует", http.StatusBadRequest)
		return
	}

	// Формируем ответ
	answer := fmt.Sprintf("Ответ на ваш вопрос '%s': 42", question)

	// Отправляем ответ клиенту
	fmt.Fprintln(w, answer)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	myObject := marshrut{
		price: 12,
	}
	// Формируем ответ
	jsonData, err := json.Marshal(myObject)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println(w, string(myObject))
	// Отправляем ответ клиенту
	fmt.Fprintln(w, string(jsonData))
}

func main() {
	// Регистрируем обработчик для пути "/"
	http.HandleFunc("/9292839", handler)
	http.HandleFunc("/111", handler2)

	// Запускаем веб-сервер на порту 8080
	fmt.Println("Сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %v\n", err)
	}
}
