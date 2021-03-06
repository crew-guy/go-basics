package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("C:/Users/ankit/Desktop/go/file-reader/test-file.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
