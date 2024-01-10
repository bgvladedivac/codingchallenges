package tests

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

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
		for _, lineToken := range lineTokens {
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
