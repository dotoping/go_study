package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("hollo.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}
