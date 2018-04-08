package main

import (
	"bufio"
	"os"
	"fmt"
	"log"
)

const (
	filePath = "./01-reader-write-file.go"
)

func main() {
	f, _ := os.Open(filePath)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	f, err := os.Create("./extras/output.txt")
	if err != nil {
		log.Fatalln("Error creating file: ", err)
	}
	defer f.Close()

	for _, str := range []string{"apple\n", "banana", "carrot"} {
		bytes, err := f.WriteString(str)
		if err != nil {
			log.Fatalln("插入失败", err)
		}
		fmt.Printf("Wrote %d bytes to file\n", bytes)
	}
}