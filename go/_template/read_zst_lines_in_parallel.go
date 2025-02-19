package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sync"
)

type argsT struct {
	threads *int
	inFile  *string
}

var (
	args argsT
	MU   = &sync.Mutex{}
	WG   = &sync.WaitGroup{}
)

func init() {
	args.threads = flag.Int("t", 16, "parallelization thread count")
	args.inFile = flag.String("i", "", "input zst file")
}

func main() {
	flag.Parse()

	queue := make(chan string, *args.threads*8)
	var count int

	for i := 0; i < *args.threads; i++ {
		WG.Add(1)
		go processLines(queue, &count)
	}

	cmd := exec.Command("zstdcat", "--long=31", *args.inFile)
	out, err := cmd.StdoutPipe()
	errExit(err)

	err = cmd.Start()
	errExit(err)

	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		queue <- scanner.Text()
	}

	close(queue)
	WG.Wait()

	fmt.Println(count)
}

func processLines(queue chan string, count *int) {
	defer WG.Done()

	for line := range queue {
		processLine(line, count)
	}
}

func processLine(line string, count *int) {
	MU.Lock()
	*count++
	MU.Unlock()
}

func errExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
