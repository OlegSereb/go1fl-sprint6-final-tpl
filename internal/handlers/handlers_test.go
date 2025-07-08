package handlers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Тест для GET-запроса к /
func TestRootHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RootHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Неверный статус: %v вместо %v", status, http.StatusOK)
	}

	if ctype := rr.Header().Get("Content-Type"); ctype != "text/html; charset=utf-8" {
		t.Errorf("Неверный Content-Type: %q вместо %q", ctype, "text/html; charset=utf-8")
	}
}

// Вспомогательная функция для создания multipart/form-data
func createFormFile(fieldName, filename string, content []byte) (*bytes.Buffer, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(fieldName, filename)
	if err != nil {
		return nil, err
	}

	_, err = part.Write(content)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Тест для POST-запроса /upload с файлом
func TestUploadHandler_WithFile(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	// Создаём файл в форме
	part, err := writer.CreateFormFile("myFile", "test.txt")
	if err != nil {
		t.Fatal(err)
	}
	part.Write([]byte("Привет"))
	writer.Close()

	req, err := http.NewRequest("POST", "/upload", body)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UploadHandler)
	handler.ServeHTTP(rr, req)

	// Проверка
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Неверный статус: %v вместо %v", status, http.StatusOK)
	}

	responseBody, _ := io.ReadAll(rr.Body)
	if !strings.Contains(string(responseBody), "Результат:") {
		t.Errorf("Ответ не содержит ожидаемый текст: %s", responseBody)
	}
}

// Тест для POST-запроса /upload с текстовым полем
func TestUploadHandler_WithText(t *testing.T) {
	form := strings.NewReader("text=Привет")
	req, err := http.NewRequest("POST", "/upload", form)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UploadHandler)

	handler.ServeHTTP(rr, req)

	// Проверяем, что запрос успешен
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Неверный статус: %v вместо %v", status, http.StatusOK)
	}

	// Проверяем, что в теле есть результат
	responseBody, _ := io.ReadAll(rr.Body)
	if !strings.Contains(string(responseBody), "Результат:") {
		t.Errorf("Ответ не содержит ожидаемый текст")
	}
}

// Тест для POST-запроса /upload без данных
func TestUploadHandler_EmptyRequest(t *testing.T) {
	req, err := http.NewRequest("POST", "/upload", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UploadHandler)

	handler.ServeHTTP(rr, req)

	// Ожидаем ошибку Bad Request
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Неверный статус: %v вместо %v", status, http.StatusBadRequest)
	}
}
