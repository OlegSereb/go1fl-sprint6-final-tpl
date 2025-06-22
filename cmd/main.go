package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	server := server.NewServer(logger)

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	log.Println("Сервер запускается...")
	if err := server.Start(); err != nil {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
