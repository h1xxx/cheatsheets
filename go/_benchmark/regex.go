package main

import (
	"fmt"
	"regexp"
	"runtime"
	"time"
)

func main() {
	SIZE :=  10000000
	SIZE =  1000

	start := time.Now()
	for i := 0; i < SIZE; i++ {
		reWSpace := regexp.MustCompile(`\s+`)
		_ = reWSpace
		//line := "xxxxxxxxxxx\txxxxxxxxxxxxxxx"
		//line = reWSpace.ReplaceAllString(line, "\t")
	}
	d1 := time.Since(start)
	fmt.Printf("%s\n", d1)

	runtime.GC()
	start = time.Now()
	reWSpace := regexp.MustCompile(`\s+`)
	for i := 0; i < SIZE; i++ {
		_ = reWSpace
		//line := "xxxxxxxxxxx\txxxxxxxxxxxxxxx"
		//line = reWSpace.ReplaceAllString(line, "\t")
	}
	d2 := time.Since(start)
	fmt.Printf("%s\n", d2)
	fmt.Printf("%%diff %2.2f%%\n",
		(float64(d1.Milliseconds()-d2.Milliseconds()))/float64(d1.Milliseconds())*100)
}

