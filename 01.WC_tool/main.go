package main

import (
	"fmt"
	"wc/handlers"
	"flag"
	"os"
)


func main() {	
	enableBytes := flag.Bool("c", false, "Get the number of bytes in a file.")
	enableLines := flag.Bool("l", false, "Get the number of lines in a file.")
	enableWords := flag.Bool("w", false, "Get the number of words in a file.")
	enableChars := flag.Bool("m", false, "Get the number of characters in a file.")
	flag.Parse()

	screenMsg := "Number of %s in the file %s is %d.\n"
	file := os.Args[len(os.Args)-1]
	
	_, err := os.Stat(file)

	if os.IsNotExist(err) {
		fmt.Printf("File does not exist.")
		panic(1)
	}


	if *enableBytes {
		bytes, _ := handlers.GetBytes(file)
		fmt.Printf(screenMsg, "bytes", file, bytes) 
	}

	if *enableLines {
		lines, _ := handlers.GetLines(file)
		fmt.Printf(screenMsg, "lines", file, lines)
	}

	if *enableWords {
		words, _ := handlers.GetWords(file)
		fmt.Printf(screenMsg, "words", file, words)
	}

	if *enableChars {
		chars, _ := handlers. GetChars(file)
		fmt.Printf(screenMsg, "chars", file, chars)
	}
}	


