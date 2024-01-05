package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
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

func main() {
	file, err := os.Open("invalid.json")	  
	check(err)
	
	curleyBrackets := []string{}
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		lineTokens := strings.Split(line, " ")
		for _, lineToken := range lineTokens {
			if lineToken == "{" {
				curleyBrackets = append(curleyBrackets, lineToken)
			}

			if lineToken == "}" {
				curleyBrackets = append(curleyBrackets, lineToken)
			}
		}
	}

	fmt.Println(curleyBrackets)
	fmt.Println(len(curleyBrackets))

	leftBraacketOccurences := countOccurences(curleyBrackets, "{")

	rightBracketOccurences := countOccurences(curleyBrackets, "}")

	fmt.Println(leftBraacketOccurences)
	fmt.Println(rightBracketOccurences)

	if leftBraacketOccurences == rightBracketOccurences {
		fmt.Println("This is a valid json file.")
	} else {
		fmt.Println("This is not a valid json file.")
	}
}
