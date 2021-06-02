package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"unicode"
	"unicode/utf8"
)

type (
	Test2 struct {
		Feild int
		Id    int
	}

	Obj struct {
		M   map[uint64]struct{}
		Mut sync.RWMutex
	}
)

var (
	m = map[int]*Test2{}
)

type First struct {
	Second
}

type Second struct {
	H string
}

func (s *Second) SayHello(name string) {
	fmt.Printf("Hello %s", name)
}

func main() {
	var c chan string
	// c = make(chan string, 1)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		c <- "vlad"
		wg.Done()
	}()
	go func() {
		fmt.Println(<-c)
		wg.Done()
	}()
	wg.Wait()
}

func Embeded() {
	f := First{}
	f.SayHello(f.H)
	fmt.Println(f.H)

}

func testUTC() {
	utc := time.Now().UTC()
	fmt.Println(utc)
	local := utc
	location, err := time.LoadLocation("Europe/Budapest")
	if err == nil {
		local = local.In(location)
	}
	fmt.Println("UTC", utc.Format("15:04"), local.Location(), local.Format("15:04"))
	local = utc
	location, err = time.LoadLocation("America/Los_Angeles")
	if err == nil {
		local = local.In(location)
	}
	fmt.Println("UTC", utc.Format("15:04"), local.Location(), local.Format("15:04"))
}

func testForWait(ch chan bool, t time.Duration, o Obj) {
	time.Sleep(t)
	ch <- true
}

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	fmt.Println("div => ", div)
	fmt.Println("before => ", round)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	fmt.Println(" after => ", round)
	newVal = round / pow
	fmt.Println("newVal => ", newVal)
	return
}

func FormMap(m map[uint64]int) {
	m[1]++
	m[2]++
}

func printPhrases() {
	phrases := []string{"make me happy", "let`s go", "I got",
		"what`s up men", "it is simple",
		"try it", "I have too long phrase now"}

	for _, phrase := range phrases {
		fmt.Printf("%30s\n", phrase)
	}
}

func round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

type test1 struct {
	Name  string   `json:"name"`  // and you can add more description
	Age   int      `json:"age"`   // and you can add more description
	Login string   `json:"login"` // and you can add more description
	Hobby []string `json:"hobby"` // and you can add more description
}

func tempDirs() []string {
	return []string{
		"first",
		"second",
		"third",
		"fourth",
		"fiveth",
		"sexth",
	}
}

func testMarshal() {
	t1 := test1{
		Name:  "vlad",
		Age:   33,
		Login: "",
		Hobby: nil,
	}

	data, err := json.MarshalIndent(t1, "", "    ")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println(string(data))

	var t2 test1

	if err := json.Unmarshal(data, &t2); err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println(t2)
}

func AppendInt(x []int, y ...int) []int {
	zlen := len(x) + len(y)
	z := make([]int, zlen)

	copy(z[:len(x)], x)
	copy(z[len(x):], y)
	return z
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func charCount(ch_out chan bool) {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	ch_out <- true
}

func controlCharCount() {
	ch_out := make(chan bool)

	go charCount(ch_out)

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		<-sigs

		fmt.Print("\rExit Ok")
		os.Exit(0)
	}()

	<-ch_out
}

func testSlice() {
	test := []int{1}

	for {
		test = AppendInt(test, test...)
		test = remove(test, 1)
		fmt.Println(test)
		time.Sleep(time.Second * 1)
	}
}
