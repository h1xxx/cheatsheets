package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

type argsT struct {
	inFile *string
}

var args argsT

func init() {
	args.inFile = flag.String("i", "", "input zst file")
}

func main() {
	flag.Parse()

	cmd := exec.Command("zstdcat", "--long=31", *args.inFile)
	out, err := cmd.StdoutPipe()
	errExit(err)

	err = cmd.Start()
	errExit(err)

	var i int

	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		line := scanner.Text()
		_ = line
		i++
	}

	fmt.Println(i)
}

func errExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
