package service

import (
	"testing"
)

func TestConvert_TextToMorse(t *testing.T) {
	input := "Привет"
	expected := ".--. .-. .. .-- . -"

	result, err := Convert(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Convert(%q) = %q; want %q", input, result, expected)
	}
}

func TestConvert_MorseToText(t *testing.T) {
	input := ".--. .-. .. .-- . -"
	expected := "ПРИВЕТ"

	result, err := Convert(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("Convert(%q) = %q; want %q", input, result, expected)
	}
}
