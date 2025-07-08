package morse

import (
	"strings"
	"testing"
)

// TestRoundTrip_RussianText проверяет, что конвертация туда и обратно работает корректно
func TestRoundTrip_RussianText(t *testing.T) {
	input := "ЕЯИЗУЦКЕГЕЧТ"

	// 1. Текст → Морзе
	morse := ToMorse(input)

	if morse == "" {
		t.Fatal("Результат конвертации в Морзе пустой")
	}

	// 2. Морзе → Текст
	text := ToText(morse)

	// 3. Проверяем совпадение (без учёта регистра)
	expected := strings.ToUpper(input)
	if text != expected {
		t.Errorf("После конвертации получено: %q; ожидалось: %q", text, expected)
	}
}

// TestRoundTrip_MixedText проверяет работу с разным регистром
func TestRoundTrip_MixedText(t *testing.T) {
	input := "еяизуцкегечт" // строчные буквы

	morse := ToMorse(input)
	if morse == "" {
		t.Fatal("Результат конвертации в Морзе пустой")
	}

	text := ToText(morse)

	expected := strings.ToUpper(input)
	if text != expected {
		t.Errorf("После конвертации получено: %q; ожидалось: %q", text, expected)
	}
}

// TestRoundTrip_WithUnsupportedChars проверяет поведение при неподдерживаемых символах
func TestRoundTrip_WithUnsupportedChars(t *testing.T) {
	input := "ЕЯИЗУЦКЕГЕЧТ!?@"

	// Используем кастомный конвертер, чтобы видеть ошибки
	customConverter := NewConverter(
		DefaultMorse,
		WithCharSeparator(" "),
		WithHandler(func(err error) string {
			if _, ok := err.(ErrNoEncoding); ok {
				return "[?]"
			}
			return ""
		}),
	)

	morse := customConverter.ToMorse(input)
	if !strings.Contains(morse, "[?]") {
		t.Errorf("Ожидалось появление [?] в результате, так как есть неподдерживаемые символы")
	}
}
