package main

import (
"fmt"
"net/http"
)

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

func main() {
// Регистрируем обработчик для пути "/"
http.HandleFunc("/", handler)

	// Запускаем веб-сервер на порту 8080
	fmt.Println("Сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %v\n", err)
	}
}
Как это работает:
Сервер слушает запросы на порту 8080.

Когда поступает запрос на путь /, вызывается функция handler.

Функция handler извлекает параметр question из URL (например, /question=Что%20такое%20жизнь).

Если параметр question отсутствует, сервер возвращает ошибку 400 Bad Request.

Если параметр question присутствует, сервер возвращает ответ в формате:

Copy
Ответ на ваш вопрос '<question>': 42
Как запустить:
Сохраните код в файл main.go.

Откройте терминал в папке с файлом.

Запустите сервер:

bash
Copy
go run main.go
Сервер будет доступен по адресу: http://localhost:8080.

Примеры запросов:
Запрос с параметром:

URL: http://localhost:8080/?question=Что%20такое%20жизнь

Ответ:

Copy
Ответ на ваш вопрос 'Что такое жизнь': 42
Запрос без параметра:

URL: http://localhost:8080/

Ответ:

Copy
Параметр 'question' отсутствует
Этот сервер минималистичен и подходит для базовых задач. Если нужно добавить больше функциональности (например, логирование, поддержку других HTTP-методов), код можно расширить. 😊

New chat
