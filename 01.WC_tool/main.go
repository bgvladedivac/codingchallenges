package main

import (
	"fmt"
	"wc/handlers"
	"flag"
)


func main() {	
	enableBytes := flag.Bool("c", false, "Get the number of bytes in a file.")
	enableLines := flag.Bool("l", false, "Get the number of lines in a file.")
	enableWords := flag.Bool("w", false, "Get the number of words in a file.")
	enableChars := flag.Bool("m", false, "Get the number of characters in a file.")
	flag.Parse()
	handlers.DummyOne()
	
	mapInput := map[bool]interface{} {
		*enableBytes: handlers.GetBytes,
		*enableLines: handlers.GetLines,
		*enableWords: handlers.GetWords, 
		*enableChars: handlers.GetChars,
	}

	for key, value := range mapInput {
		fmt.Printf("Key -> %v, Value -> %v", key, value)
	}
}
