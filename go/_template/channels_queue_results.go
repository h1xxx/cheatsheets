package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	workers := 16
	queue := make(chan int, workers*2)
	results := make(chan int, workers*2)
	wg := new(sync.WaitGroup)

	for w := 1; w <= workers; w++ {
		wg.Add(1)
		go increment(queue, results, wg)
	}

	go func() {
		for i := 1; i <= 100; i++ {
			queue <- i
		}
		close(queue)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for i := range results {
		fmt.Println(i)
	}
}

func increment(queue <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range queue {
		for j := 1; j <= 1000000000; j++ {
			i++
		}
		results <- i
	}
}

func errExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
