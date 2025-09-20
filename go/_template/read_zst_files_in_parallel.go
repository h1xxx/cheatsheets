package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sync"

	fp "path/filepath"
)

type argsT struct {
	threads *int
	inFiles *string
}

var args argsT

func init() {
	args.threads = flag.Int("t", 1, "parallelization thread count")
	args.inFiles = flag.String("i", "", "glob with input zst files")
}

func main() {
	flag.Parse()

	files, err := fp.Glob(*args.inFiles)
	errExit(err)

	queue := make(chan string, *args.threads*4)
	wg := new(sync.WaitGroup)

	for i := 0; i < *args.threads; i++ {
		wg.Add(1)
		go processFiles(queue, wg)
	}

	for _, file := range files {
		queue <- file
	}

	close(queue)
	wg.Wait()
}

func processFiles(queue chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for file := range queue {
		processFile(file)
	}
}

func processFile(file string) {
	cmd := exec.Command("zstdcat", "--long=31", file)
	out, err := cmd.StdoutPipe()
	errExit(err)

	cmd.Stderr = os.Stderr

	err = cmd.Start()
	errExit(err)

	var count int

	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		line := scanner.Text()
		_ = line
		count++
	}

	err = scanner.Err()
        if err != nil {
                return fmt.Errorf("scanner: %w", err)
        }

        err = cmd.Wait()
        if err != nil {
                return fmt.Errorf("command wait: %w", err)
        }

	fmt.Println(count)
}

func errExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
