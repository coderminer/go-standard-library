package main

import (
	"log"
	"os"
)

func main() {
	//简单的打开文件
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	
	file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}
