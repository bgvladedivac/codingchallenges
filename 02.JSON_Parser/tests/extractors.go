package tests

import "fmt"

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

func ExtractJsonDocLines(fileName string) ([]JSONDocLine, error) {

}
