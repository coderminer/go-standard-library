package main

import (
	"io/ioutil"
	"log"
)

func main() {
	err := ioutil.WriteFile("test.text", []byte("测试快速写入功能!"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
