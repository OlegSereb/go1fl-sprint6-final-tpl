package morse

import (
	"testing"
)

// Тест для функции ToMorse — преобразование текста в код Морзе
func TestToMorse(t *testing.T) {
	input := "Привет"
	expected := ".--. .-. .. .-- . -"

	result := ToMorse(input)

	if result != expected {
		t.Errorf("ToMorse(%q) = %q; want %q", input, result, expected)
	}
}

// Тест для функции ToText — преобразование кода Морзе обратно в текст
func TestToText(t *testing.T) {
	input := ".--. .-. .. .-- . -"
	expected := "ПРИВЕТ"

	result := ToText(input)

	if result != expected {
		t.Errorf("ToText(%q) = %q; want %q", input, result, expected)
	}
}
