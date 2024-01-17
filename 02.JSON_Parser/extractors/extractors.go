package extractors

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type JSONDocument struct {
	jsonLines []JSONDocLine
}

type JSONDocLine struct {
	key   string
	value interface{}
}

func CreateNewJSONDocLine(key string, value interface{}) *JSONDocLine {
	jsDocLine := JSONDocLine{
		key:   key,
		value: value,
	}
	return &jsDocLine
}

func (jsDocLine *JSONDocLine) String() string {
	lineStr := fmt.Sprintf("I am a JSON Doc line with the content of key -> %s, Value -> %v.\n", jsDocLine.key, jsDocLine.value)
	return lineStr
}

// Extract basic key, value pair lines
func ExtractJsonDocLines(fileName string) ([]*JSONDocLine, error) {
	result := []*JSONDocLine{}

	file, err := os.Open(fileName)
	if err != nil {
		errMsg := fmt.Sprintf("Error openinig file %s", fileName)
		return nil, errors.New(errMsg)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		clearLine := strings.ReplaceAll(line, " ", "")
		if len(clearLine) > 1 {
			lineTokens := strings.Split(strings.ReplaceAll(clearLine, ",", ""), ":")
			jsonDocLine := CreateNewJSONDocLine(lineTokens[0], lineTokens[1])
			result = append(result, jsonDocLine)
		}
	}

	return result, nil
}
