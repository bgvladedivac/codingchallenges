package extractors

import (
	"fmt"
	"testing"
)

func TestExtractingJsonDocLine(t *testing.T) {
	file := "./data/tests.json"
	result, err := ExtractJsonDocLines(file)

	if err != nil {
		t.Errorf("Testing filename %s not found.", file)
	}

	fmt.Println("This is my result")
	fmt.Println(result)

	for key, value := range result {
		fmt.Println(key, " -> ", value)
	}

	testPairs := []struct {
		key   string
		value interface{}
	}{
		{"key1", true},
		{"key2", false},
		{"key3", nil},
		{"key4", "value"},
		{"key5", 101},
	}

	for indx, testKey := range testPairs {
		`
	}

}
