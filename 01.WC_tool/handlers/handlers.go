package handlers

import (
	"os"
	"bufio"
	"fmt"
	"errors"
	"regexp"
	"reflect"
	"io/ioutil"
)


func GetWords(fileName string) (int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		errorMsg := fmt.Sprintf("File %s was not found.", fileName)
		return -1, errors.New(errorMsg)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordRegex := regexp.MustCompile(`\b\w+\b`)
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		matches := wordRegex.FindAllString(line, -1)
		counter += len(matches)
	}
	
	return counter, nil

}

func GetLines(fileName string) (int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		errorMsg := fmt.Sprintf("File %s was not found.", fileName)
		return -1, errors.New(errorMsg)
	}

	defer file.Close()

	numberOfLines := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		scanner.Text()
		numberOfLines += 1
	}

	return numberOfLines, nil
}

func GetBytes(fileName string) (int, error) {
	fileInfo, err := os.Stat(fileName)

	if err != nil {
		errorMsg := fmt.Sprintf("File %s was not found.", fileName)
		return -1, errors.New(errorMsg)
	}

	return int(fileInfo.Size()), nil
}

func GetChars(fileName string) (int, error) {
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		errMsg := fmt.Sprintf("Filename %s was not found.", fileName)	
		return -1, errors.New(errMsg)
	}

	
	return len(content), nil
}
