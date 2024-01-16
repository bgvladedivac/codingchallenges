package main

import (
	"bufio"
	"fmt"
	"jsonparser/utils"
	"os"
	"strings"
)

func main() {
	fileNames, _ := utils.GetFiles("./", ".json")
	for _, fileName := range fileNames {
		curleyBracketsMatching, err := utils.CheckIfCurlyBracketsAreMatching(fileName)

		if err != nil {
			fmt.Println("Error opening the file.")
		}

		if curleyBracketsMatching {
			file, _ := os.Open(fileName)
			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				line := scanner.Text()
				clearLine := strings.ReplaceAll(line, " ", "")
				if len(clearLine) > 1 {
					lineTokens := strings.Split(strings.ReplaceAll(clearLine, ",", ""), ":")
					fmt.Println(lineTokens)
					jsonLineDoc := CreateNewJSONDocumentLine(lineTokens[0], lineTokens[1])
					fmt.Println(jsonLineDoc)

				}
			}

		}
	}

}
