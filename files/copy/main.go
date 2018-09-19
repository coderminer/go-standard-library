package main

import (
	"io"
	"log"
	"os"
)

func main() {
	//打开原文件
	originalFile, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	//创建新文件
	newFile, err := os.Create("test_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	//文件复制
	bytes, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytes)

	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
