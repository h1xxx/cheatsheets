package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"syscall"
	"time"

	str "strings"
)

var baseDate = time.Date(2000, time.Month(1), 1, 0, 0, 0, 0, time.UTC)

var reHttp = regexp.MustCompile(`^https?`)
var reNumbers = regexp.MustCompile(`^[0-9.,-/–$£:'&#\\%@*+=÷×!~^]*$`)
var rePunct = regexp.MustCompile(`[.,:]`)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	for i, el := range arr {
		fmt.Println(len(arr), i, el)
		arr = append(arr, 666)
		time.Sleep(1 * time.Second)
	}

	fmt.Println(len(arr), arr[8])
	os.Exit(0)

	hour := time.Now().UTC().Hour()
	logFile := "get_peers_2025-01-03.log.zst"
	re := regexp.MustCompile(`20[0-9][0-9]-[0-9][0-9]-[0-9][0-9]`)
	partString := re.FindString(logFile)
	fmt.Println("zzz:", partString)

	i := 5
	td := time.Duration(i) * time.Second
	fmt.Println(td)
	os.Exit(0)

	for i := 0; i < 1; i++ {
		s := "test1"
		dt := time.Now().UTC()
		//dtStr := dt.UTC().Format("2006-01-02 15:04:05")

		printRef(&s)
		hexDec, err := hex.DecodeString("0000616161")
		errExit(err)
		s1 := string(hexDec)
		fmt.Println(s1)
		fmt.Println(hexDec[0] == byte(0))
		fmt.Println(hexDec[1] == byte(0))

		if hour != dt.UTC().Hour() {
			fmt.Println("new hour")
		}
	}

	os.Exit(0)

	ip := net.ParseIP("192.168.0.0")
	fmt.Println("multicast:", ip.IsInterfaceLocalMulticast())

	ids := readKnownIDs()
	fmt.Println(len(ids))
	_ = ids

	fmt.Println(isValidEmail("kot@tlen.org"))
	fmt.Println(isValidEmail("kot@tlen-org.pl"))
	fmt.Println(isValidEmail("kot_sf+shit@tlen-org.zinc.pl"))
	fmt.Println(isValidEmail("kot_sf+shit@tlen-org.zinc.pl.so"))

	fmt.Println(isValidEmail("kot_sf+shit@tlen-org.zinc.pl.so.so"))
	fmt.Println(isValidEmail("kot@tlenorg"))
	fmt.Println(isValidEmail("kot@tlen-org"))

	time.Sleep(100 * time.Second)
}

func printRef(s *string) {
	fmt.Println(*s)
}

func readKnownIDs() map[[8]byte]struct{} {
	file := "live_id.lst"
	fd, err := os.Open(file)
	errExit(err)

	knownIDs := make(map[[8]byte]struct{})

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		hashID := hashIDStrToByte(line[:16])
		knownIDs[hashID] = struct{}{}
	}

	fd.Close()

	return knownIDs
}

func hashIDStrToByte(s string) [8]byte {
	var hashID [8]byte

	hashIDSlice, err := hex.DecodeString(s)
	errExit(err)

	copy(hashID[:], hashIDSlice[:8])

	return hashID
}

func main_old() {
	fmt.Println(DateStrToUint("2023-12-24"))
	matched := reHttp.MatchString("httxps")
	fmt.Println(matched)
	matched = reHttp.MatchString("http")
	fmt.Println(matched)

	fmt.Println("---------")

	matched = reNumbers.MatchString("h,.ttp0")
	fmt.Println(matched)

	matched = reNumbers.MatchString("^3445*!0&")
	fmt.Println(matched)

	matched = reNumbers.MatchString("34.45,0")
	fmt.Println(matched)

	matched = reNumbers.MatchString("3:4.%'4-/50")
	fmt.Println(matched)

	matched = reNumbers.MatchString("34,450")
	fmt.Println(matched)

	matched = reNumbers.MatchString(`3=3=2+\1`)
	fmt.Println(matched)

	matched = rePunct.MatchString("34:450")
	fmt.Println(matched)

	for i := 0; i < 10000; i++ {
		TimeStrToByteA("2023-12-24 13:59:01")
	}

	v1, err := TimeStrToByteA("2023-12-24 13:59:01")
	if err != nil {
		fmt.Println(err)
	}

	v2, err := TimeStrToByteA2("2023-12-24 13:59:01")
	if err != nil {
		fmt.Println(err)
	}
	var slice []string
	for i := 0; i < 1000000; i++ {
		slice = append(slice, fmt.Sprintf("zzzzzzzzzzzzz%d", i))
	}

	fmt.Printf("%x\n", v1)
	fmt.Printf("%x\n", v2)

	bs := make([]byte, 2)
	bs2 := []byte{}
	fmt.Println(bs)
	fmt.Println(bs2)

	fmt.Printf("pid: %d\n", syscall.Getpid())
	fmt.Printf("mem: %d\n", getProcessMem())

	time.Sleep(0 * time.Second)
	os.Exit(0)
}

func isValidEmail(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._+\-]+@[a-z0-9\-]+(\.[a-z]{2,4}){1,3}$`)
	return emailRegex.MatchString(e)
}

func getProcessMem() int {
	file := fmt.Sprintf("/proc/%d/smaps", syscall.Getpid())
	b, _ := os.ReadFile(file)
	lines := str.Split(string(b), "\n")

	var mem int
	for _, line := range lines {
		if str.HasPrefix(line, "Rss:") {
			fields := str.Fields(line)
			rss, _ := strconv.Atoi(fields[1])
			mem += rss
		}
	}
	return mem
}

func TimeStrToByteA(dateStr string) ([3]byte, error) {
	var timeA [3]byte

	datetime, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		return timeA, err
	}

	bs := make([]byte, 2)

	num := uint16(datetime.Hour())
	binary.LittleEndian.PutUint16(bs, num)
	timeA[0] = bs[0]

	num = uint16(datetime.Minute())
	binary.LittleEndian.PutUint16(bs, num)
	timeA[1] = bs[0]

	num = uint16(datetime.Second())
	binary.LittleEndian.PutUint16(bs, num)
	timeA[2] = bs[0]

	return timeA, nil
}

func TimeStrToByteA2(dateStr string) ([3]byte, error) {
	var timeA [3]byte
	datetime, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		return timeA, err
	}

	buf := new(bytes.Buffer)

	num := uint16(datetime.Hour())
	err = binary.Write(buf, binary.LittleEndian, num)
	if err != nil {
		return timeA, err
	}
	timeA[0] = buf.Bytes()[0]

	buf = new(bytes.Buffer)
	num = uint16(datetime.Minute())
	err = binary.Write(buf, binary.LittleEndian, num)
	if err != nil {
		return timeA, err
	}
	timeA[1] = buf.Bytes()[0]

	buf = new(bytes.Buffer)
	num = uint16(datetime.Second())
	err = binary.Write(buf, binary.LittleEndian, num)
	if err != nil {
		return timeA, err
	}
	timeA[2] = buf.Bytes()[0]

	return timeA, nil
}

func DateStrToUint(dateStr string) (uint16, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return 0, err
	}

	return DateToUint(date), nil
}

func DateToUint(datetime time.Time) uint16 {
	return uint16(datetime.Sub(baseDate).Hours() / 24)
}

func errExit(err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
