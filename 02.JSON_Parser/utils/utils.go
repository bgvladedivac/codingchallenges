package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// get the number of items an iteam is found in a list
func CountOccurences(s []string, item string) int {
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
func CheckIfCurlyBracketsAreMatching(fileName string) (bool, error) {
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
		clearLine := strings.ReplaceAll(line, " ", "")
		for _, char := range clearLine {
			c := string(char)
			if strings.Compare(c, leftCurley) == 0 {
				curleyBrackets = append(curleyBrackets, c)
			}

			if strings.Compare(c, rightCurley) == 0 {
				curleyBrackets = append(curleyBrackets, c)
			}
		}
	}

	leftCurleyCounter := CountOccurences(curleyBrackets, leftCurley)
	rightCurleyCounter := CountOccurences(curleyBrackets, rightCurley)
	// consider the file to be valid json if '{' count is equal to '}'
	return leftCurleyCounter == rightCurleyCounter, nil
}

// picks up all files with a specific extension, useful for easier testing
func GetFiles(dir, extens string) ([]string, error) {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		errMsg := fmt.Sprintf("Error while reading the directory %s", dir)
		return nil, errors.New(errMsg)
	}

	files := []string{}
	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() && strings.HasSuffix(dirEntry.Name(), extens) {
			files = append(files, dirEntry.Name())

		}
	}

	return files, nil
}
