package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	str "strings"
)

func main() {
	lines := readLines("file1.txt")
	fmt.Println(lines)
}

func readLines(file string) []string {
	fd, err := os.Open(file)
	errExit(err, "can't open file")
	scanner := bufio.NewScanner(fd)

	//buf := make([]byte, 0, 64*1024)
	//scanner.Buffer(buf, 512*1024*1024)

	var lines []string
	for scanner.Scan() {
		line := str.TrimSpace(scanner.Text())

		switch {
		case line == "":
			continue
		case str.HasPrefix(line, "#"):
			continue
		}

		lines = append(lines, line)
	}
	err = scanner.Err()
	errExit(err)

	return lines
}

func errExit(err error, msg string) {
	if err != nil {
		log.Println("\n * " + msg)
		os.Exit(1)
	}
}
