package tests

import (
	"testing"
	"wc/handlers"
)

const testFile string = "dummy.txt"

func TestGetLines(t *testing.T) {
	expected := 4 
	got, _ := handlers.GetLines(testFile)

	if expected != got {
		t.Errorf("Number of lines is not matching.\nExpected %d, Got %dl", expected, got)
	}
}

func TestGetBytes(t *testing.T) {
	expected := 1
	got, _ := handlers.GetBytes(testFile)

	if expected > got {
		t.Errorf("There are no bytes in this file anon.")
	}
}

func TestGetWords(t *testing.T) {
	expected := 10
	got, _ := handlers.GetWords(testFile)

	if expected != got {
		t.Errorf("Number of words are not matching.\nExpected: %d, Got: %d", expected, got)
	}
}

func TestGetChars(t *testing.T) {
	expected := 51
	got, _ := handlers.GetChars(testFile)

	if expected != got {
		t.Errorf("Number of chars are not matching.\nExpected: %d, Got: %d", expected, got)
	}
	
}
