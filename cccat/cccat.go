package main

import (
	"io"
	"log"
	"os"
)

func ReadFile(writer io.Writer, path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	writer.Write(data)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments to cccat")
	}
	ReadFile(os.Stdout, os.Args[1])
}
