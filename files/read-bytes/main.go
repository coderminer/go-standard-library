package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes := make([]byte, 16)
	br, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("number of bytes read: %d\n", br)
	log.Printf("Data read: %s\n", bytes)
}
