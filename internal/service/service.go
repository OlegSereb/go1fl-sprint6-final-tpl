package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert автоматически определяет тип данных: текст или код Морзе.
// Если содержит только точки и тире — считаем, что это код Морзе.
func Convert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)

	// Если строка содержит только '.', '-', ' ' — скорее всего, это код Морзе
	if isMorse(trimmed) {
		return morse.ToText(trimmed), nil
	}

	// Иначе преобразуем в код Морзе
	return morse.ToMorse(trimmed), nil
}

// isMorse проверяет, является ли строка кодом Морзе
func isMorse(s string) bool {
	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' && r != '\n' {
			return false
		}
	}
	return len(s) > 0
}
