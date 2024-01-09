package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"errors"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func countOccurences(s []string, item string) int {
	occurences := 0

	for _, value := range s {
		if strings.Compare(value, item) == 0 {
			occurences += 1
		}
	}

	return occurences
}


// checks if the number of left curly bracket '{'
// equals the number of right curly bracket '}'
func checkIfCurlyBracketsAreMatching(fileName string) (bool, error) {
	file, err := os.Open(fileName)
	if err != nil {
		errMsg := fmt.Sprintf("Filename %s does not exist.", fileName)
		
		return false, errors.New(errMsg)
	}

	defer file.Close()

	curleyBrackets := []string{}
	leftCurley := "{"
	rightCurley := "}"
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		lineTokens := strings.Split(line, " ")
		for _ , lineToken := range lineTokens {
			if lineToken == leftCurley {
				curleyBrackets = append(curleyBrackets, leftCurley)
			}

			if lineToken == rightCurley {
				curleyBrackets = append(curleyBrackets, rightCurley)
			}
		}
	}
	
	leftCurleyCounter := countOccurences(curleyBrackets, leftCurley)
	rightCurleyCounter := countOccurences(curleyBrackets, rightCurley)

	// consider the file to be valid json if '{' count is equal to '}'
	return leftCurleyCounter == rightCurleyCounter, nil

}


// checks if each line has a double dot and a previous quote
// similiar to ": 
func checkIfDoubleDotAndPreviousQuoteArePresentPerLine(fileName string) (bool, error) {
	file, err := os.Open(fileName)

	if err != nil {
		errMsg := fmt.Sprintf("Filename %s does not exist.", fileName)
		return false, errors.New(errMsg)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		validLine := false

		for index, char := range line {
			c := string(char)
			if strings.Compare(c, ":") == 0 {
				if string(line[index-1]) == "\"" {
					validLine = true	
				}
			}
		}

		if !validLine  {
			return false, nil
		}
	}

	return true, nil
}

// declare a method to pick up all json files and store them in the slice
func GetFiles(directory string, extension string) ([]string, error) {
	dirEntriers, err := os.ReadDir(directory)
	if err != nil {
		errMsg := fmt.Sprintf("Error while reading the directory %s", directory)
		return nil, errors.New(errMsg)
	}

	var files []string
	
	for _, dirEntry := range dirEntriers {
		if !dirEntry.IsDir() && strings.HasSuffix(dirEntry.Name(), extension) {
			files = append(files, dirEntry.Name())
		}
	}

	return files, nil
}


// declare a method to squuze all empty lines in the json file.

func main() {
	fileNames, _ := GetFiles("./", ".json")

	for _, fileName := range fileNames {
		fmt.Println("Result for file %s are as following:", fileName)
		fmt.Println(checkIfCurlyBracketsAreMatching(fileName))
		fmt.Println(checkIfDoubleDotAndPreviousQuoteArePresentPerLine(fileName))
		fmt.Println()
	}


}
