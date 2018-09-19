package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("empty.txt")
	if err != nil {
		log.Fatal("create file err", err)
	}
	log.Println(file)
	file.Close()
}
