package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("test/test")
	errExit(err, "can't open file")
	scanner := bufio.NewScanner(f)

	//buf := make([]byte, 0, 64*1024)
	//scanner.Buffer(buf, 512*1024*1024)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Join(strings.Fields(line), "\t")
		switch {
		case strings.HasPrefix(line, "name"):
			line = convName(line)
		}
		fmt.Println(line)
	}
}

func convName(line string) string {
	l := strings.Replace(line, "name\t", "N:", -1)
	return l
}

func errExit(err error, msg string) {
	if err != nil {
		log.Println("\n * " + msg)
		log.Fatal(err)
	}
}
