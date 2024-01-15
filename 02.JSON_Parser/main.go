package main

import (
	"bufio"
	"fmt"
	"jsonparser/utils"
	"os"
	"strings"
)

type JSONDocument struct {
	jsonLines []JSONDocumentLine
}

type JSONDocumentLine struct {
	key   string
	value any
}

func CreateNewJSONDocumentLine(key string, value any) *JSONDocumentLine {
	jsDocLine := JSONDocumentLine{
		key:   key,
		value: value,
	}
	return &jsDocLine
}

func (jsDocLine *JSONDocumentLine) String() string {
	lineStr := fmt.Sprintf("I am a JSON Document Line with Content of Key -> %s, Value -> %v.\n", jsDocLine.key, jsDocLine.value)
	return lineStr
}

func main() {
	fileNames, _ := utils.GetFiles("./", ".json")
	fmt.Println("Length of file names -> ", len(fileNames))
	for _, fileName := range fileNames {
		curleyBracketsMatching, err := utils.CheckIfCurlyBracketsAreMatching(fileName)

		if err != nil {
			fmt.Println("Error opening the file.")
		}

		if curleyBracketsMatching {
			file, _ := os.Open(fileName)
			scanner := bufio.NewScanner(file)

			key := ""
			value := ""
			for scanner.Scan() {
				line := scanner.Text()
				clearLine := strings.ReplaceAll(line, " ", "")
				splittedClearLine := strings.Split(clearLine, ":")
				fmt.Println("\n")

				for index, token := range splittedClearLine {
					if index == 0 {
						key = token
					}

					if index == 1 {
						value = token
					}
				}

			}

			jsonDocLine := CreateNewJSONDocumentLine(key, value)
			fmt.Println(jsonDocLine)

		}
	}

}
