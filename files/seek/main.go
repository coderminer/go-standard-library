package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var offset int64 = 5

	var whence int = 0
	newPos, err := file.Seek(offset, whence)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Just moved to 5:", newPos)

	newPos, err = file.Seek(-2, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Just moved back two:", newPos)

	currentPos, err := file.Seek(0, 1)
	fmt.Println("current pos:", currentPos)

	newPos, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("position after seek 0,0:", newPos)
}
