package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/static"
)

// RootHandler отправляет HTML-форму
func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Обработка запроса к /")

	fileServer := http.FileServer(http.FS(static.Files))
	fileServer.ServeHTTP(w, r)
}

// UploadHandler принимает файл, конвертирует и сохраняет результат
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Ограничиваем размер загружаемого контента
	r.ParseMultipartForm(10 << 20) // 10 MB

	var content []byte
	var err error

	// Пробуем получить файл
	file, handler, err := r.FormFile("myFile")
	if err == nil {
		defer file.Close()
		log.Printf("Загружен файл: %s", handler.Filename)

		content, err = io.ReadAll(file)
		if err != nil {
			log.Printf("Ошибка чтения файла: %v", err)
			http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
			return
		}
	} else {
		// Если файла нет, пробуем получить текст из поля "text"
		text := r.FormValue("text")
		if text == "" {
			http.Error(w, "Не передан ни файл, ни текст", http.StatusBadRequest)
			return
		}
		content = []byte(text)
	}

	// Конвертируем
	result, err := service.Convert(string(content))
	if err != nil {
		log.Printf("Ошибка конвертации: %v", err)
		http.Error(w, "Ошибка конвертации", http.StatusInternalServerError)
		return
	}

	// Генерируем имя файла
	filename := "converted_" + time.Now().UTC().Format("20060102_150405") + ".txt"

	// Создаём и записываем новый файл
	newFile, err := os.Create(filename)
	if err != nil {
		log.Printf("Ошибка создания файла: %v", err)
		http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = newFile.WriteString(result)
	if err != nil {
		log.Printf("Ошибка записи файла: %v", err)
		http.Error(w, "Ошибка записи файла", http.StatusInternalServerError)
		return
	}

	// Возвращаем результат
	log.Printf("Исходный текст: %s", string(content))
	log.Printf("Результат конвертации: %s", result)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "Исходный текст: %s\n\nРезультат:\n%s\n\nСохранён как: %s", string(content), result, filename)

	// Пример использования конвертера
	fmt.Println("Пример из версии с улучшенным кодом!")
}
