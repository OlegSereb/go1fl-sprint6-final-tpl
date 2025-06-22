package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// RootHandler отправляет HTML-форму
func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Обработка запроса к /")
	wd, _ := os.Getwd()
	log.Printf("Текущая директория: %s", wd)

	http.ServeFile(w, r, "index.html")
}

// UploadHandler принимает файл, конвертирует и сохраняет результат
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// Ограничиваем размер загружаемого файла
	r.ParseMultipartForm(10 << 20) // 10 MB

	// Получаем файл из формы
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Printf("Ошибка при получении файла: %v", err)
		http.Error(w, "Ошибка при получении файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	log.Printf("Загружен файл: %s", handler.Filename)

	// Читаем содержимое файла
	content, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Ошибка чтения файла: %v", err)
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	// Конвертируем
	result, err := service.Convert(string(content))
	if err != nil {
		log.Printf("Ошибка конвертации: %v", err)
		http.Error(w, "Ошибка конвертации", http.StatusInternalServerError)
		return
	}

	// Генерируем безопасное имя файла
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "Результат:\n%s\n\nСохранён как: %s", result, filename)
}
