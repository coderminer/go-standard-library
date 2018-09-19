package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes := []byte("测试写入功能！")
	bw, err := file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bw)
}
