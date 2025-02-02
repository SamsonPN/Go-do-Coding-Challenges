package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func ReadStdIn(writer io.Writer) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		writer.Write(scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("[cccat:ReadStdIn()]:", err)
	}
}

func ReadFile(writer io.Writer, path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("[cccat:ReadFile()]:", err)
	}
	writer.Write(data)
}

func ConcatFiles(writer io.Writer, path1, path2 string) {
	ReadFile(writer, path1)
	ReadFile(writer, path2)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("[cccat:Main()]: Not enough arguments to cccat")
	}

	if os.Args[1] == "-" {
		ReadStdIn(os.Stdout)
	} else if len(os.Args) == 3 {
		ConcatFiles(os.Stdout, os.Args[1], os.Args[2])
	} else {
		ReadFile(os.Stdout, os.Args[1])
	}
}
