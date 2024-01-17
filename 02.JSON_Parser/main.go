package main

import (
	"fmt"
	"jsonparser/utils"
)

func main() {
	fileNames, _ := utils.GetFiles("./", ".json")
	for _, fileName := range fileNames {
		curleyBracketsMatching, err := utils.CheckIfCurlyBracketsAreMatching(fileName)

		if err != nil {
			fmt.Println("Error opening the file.")
		}

		if curleyBracketsMatching {

		}
	}

}
