package main

import (
	"log"
	"os"
	"time"
)

func main() {
	//改变权限
	err := os.Chmod("test.txt", 0777)
	if err != nil {
		log.Println(err)
	}

	//改变所有权 适用于linux
	err = os.Chown("test.txt", os.Getuid(), os.Getegid())
	if err != nil {
		log.Println(err)
	}

	//改变时间戳
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime)
	if err != nil {
		log.Println(err)
	}
}
