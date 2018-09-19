package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("main.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Printf("System interface type:%T\v\n", fileInfo.Sys())
	fmt.Printf("System info:%+v\n\n", fileInfo.Sys())
}
