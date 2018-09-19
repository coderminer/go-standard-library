package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("empty.txt")
	if err != nil {
		log.Fatal(err)
	}
}
