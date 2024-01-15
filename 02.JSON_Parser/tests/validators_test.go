package test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"jsonparser/utils"
	"os"
	"testing"
)

const dir = "/tmp/"

type fileTest struct {
	fileName string
	content  string
	expected bool
}

func TestCountOccurences(t *testing.T) {
	tests := []struct {
		sequence         []string
		target           string
		expecteddCounter int
	}{
		{[]string{"a", "b", "c"}, "a", 1},
		{[]string{"a", "a", "b", "c"}, "a", 2},
		{[]string{"a", "b", "c", "c", "c"}, "c", 3},
	}

	for _, test := range tests {
		got := utils.CountOccurences(test.sequence, test.target)
		if got != test.expecteddCounter {
			t.Errorf("Count occurences test is failing, Expected: %d, Got: %d", test.expecteddCounter, got)
		}
	}
}

func CreateTmpFile(dir, fileName string) (*os.File, error) {
	tmpFile, err := ioutil.TempFile(dir, fileName)

	if err != nil {
		errMsg := fmt.Sprintf("Error creating the tmp file inside dir %s, with name %s", dir, fileName)
		return nil, errors.New(errMsg)
	}

	return tmpFile, nil
}

func CloseFileHandler(file *os.File) {
	defer os.Remove(file.Name())
}

func TestCurlyBracketsMatching(t *testing.T) {
	tests := []fileTest{
		{"matching-1.json", "{\n}", true},
		{"matching-2.json", "{\"name\":\"gosho\"}", true},
	}

	for _, test := range tests {
		file, err := CreateTmpFile(dir, test.fileName)
		if err != nil {
			t.Errorf("Creating the tmp file failed.")
		}

		_, err = file.WriteString(test.content)

		if err != nil {
			t.Errorf("Writing content to tmp file failed.")
		}

		result, _ := utils.CheckIfCurlyBracketsAreMatching(file.Name())
		if result != test.expected {
			t.Errorf("Matching brackets is failing, Expected: %t, Got %t", test.expected, result)
		}

		CloseFileHandler(file)
	}
}

func TestCurlyBracketsAreNotMatching(t *testing.T) {
	tests := []fileTest{
		{"not-matching-1.json", "{", false},
		{"not-matching-2.json", "{\"name\": \"gosho\"", false},
	}

	for _, test := range tests {
		file, err := CreateTmpFile(dir, test.fileName)
		if err != nil {
			t.Errorf("Creating the tmp file failed.")
		}

		_, err = file.WriteString(test.content)
		if err != nil {
			t.Errorf("Writing content to tmp file failed.")
		}

		result, _ := utils.CheckIfCurlyBracketsAreMatching(file.Name())
		if result != test.expected {
			fmt.Println("This is the test for which it's a true")
			fmt.Println(test)
			t.Errorf("Matching brackets is failing, Expected: %t, Got %t", test.expected, result)
		}

		CloseFileHandler(file)
	}
}
