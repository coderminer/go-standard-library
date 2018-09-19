package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffer := bufio.NewWriter(file)
	bw, err := buffer.Write([]byte{65, 66, 67})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written:%d\n", bw)
	bw, err = buffer.WriteString("\n写入字符串")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bw)

	unFlushedBuffer := buffer.Buffered()
	log.Printf("Bytes buffered:%d\n", unFlushedBuffer)

	ba := buffer.Available()
	log.Printf("Available buffer: %d\n", ba)

	buffer.Flush()

	buffer.Reset(buffer)

	ba = buffer.Available()
	log.Printf("Availabled buffer:%d\n", ba)

	buffer = bufio.NewWriterSize(buffer, 8000)
	ba = buffer.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", ba)

}
