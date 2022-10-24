package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

type sT struct {
	ID       int
	By       string
	Score    int
	Comments int
	TimeI    int64
	Title    string
	Url      string
	Type     string
	Domain   string
	ScoreAvg int
	Hours    int
	sS0      string
	sS1      string
	sS2      string
	sS3      string
	sS4      string
	sS5      string
	sS6      string
	sS7      string
	sS8      string
	sS9      string
	sS10     string
	sS11     string
	sS12     string
	sS13     string
	sS14     string
	sS15     string
	sS16     string
	sS17     string
	sS18     string
	sS19     string
}

func main() {
	SIZE := 100000000
	SIZE =  10000000

	var s sT
	start := time.Now()
	for i := 0; i < SIZE; i++ {
		s = updateS1(s)
	}
	d1 := time.Since(start)
	fmt.Printf("%s    %d\n", d1, s.ID)

	s = sT{}
	runtime.GC()
	start = time.Now()
	for i := 0; i < SIZE; i++ {
		updateS2(&s)
	}
	d2 := time.Since(start)
	fmt.Printf("%s    %d\n", d2, s.ID)
	fmt.Printf("%%diff %2.2f%%\n",
		(float64(d1.Milliseconds()-d2.Milliseconds()))/float64(d1.Milliseconds())*100)
}

func updateS1(s sT) sT {
	s.ID += 1
	s.TimeI += 7
	s.Score += 4
	s.Comments += 8
	s.Hours += 5
	s.By = strconv.Itoa(s.ID)
	s.Url = strconv.FormatInt(int64(s.TimeI), 10)
	s.Domain = strconv.Itoa(s.Hours)
	s.sS0 = s.By + "lksdhfghjghsadohjkfhkjdsfnkjgkjdfgjfgfdhkjghslkf"
	s.sS1 = s.Url + "lksdhfghj3hsadohjkf1kjdsfnkjgkjdfgjfgfdhkjghslkf"
	s.sS2 = s.Domain + "1ksdhfghj3hsadohjkf1kjdsfnkj8kjdfgjfgfdhkjghslkf"
	s.sS10 = s.By + "l1sdhfghjghsadohjkfhkjdsfnkjgkjdfgjfgfdhkjghslkf"
	s.sS11 = s.Url + "lksdhfghj3hsadohj3f1kjdsfnkjgkjdfgjfgfdhkjghslkf"
	s.sS12 = s.Domain + "1ksdhfghj3hsadohjkf1kjdsfnkj88jdfgjfgfd8kjghslkf"

	return s
}

func updateS2(s *sT) {
	s.ID += 1
	s.TimeI += 7
	s.Score += 4
	s.Comments += 8
	s.Hours += 5
	s.By = strconv.Itoa(s.ID)
	s.Url = strconv.FormatInt(int64(s.TimeI), 10)
	s.Domain = strconv.Itoa(s.Hours)
	s.sS0 = s.By + "lksdhfghjghsadohjkfhkjdsfnkjgkjdfgjfgfdhkjghslkf"
	s.sS1 = s.Url + "lksdhfghj3hsadohjkf1kjdsfnkjgkjdfgjfgfdhkjghslkf"
	s.sS2 = s.Domain + "1ksdhfghj3hsadohjkf1kjdsfnkj8kjdfgjfgfdhkjghslkf"
	s.sS10 = s.By + "l1sdhfghjghsadohjkfhkjdsfnkjgkjdfgjfgfdhkjghslkf"
	s.sS11 = s.Url + "lksdhfghj3hsadohj3f1kjdsfnkjgkjdfgjfgfdhkjghslkf"
	s.sS12 = s.Domain + "1ksdhfghj3hsadohjkf1kjdsfnkj88jdfgjfgfd8kjghslkf"
}
