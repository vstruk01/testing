package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/vstruk01/testing/workerpool"
)

func main() {
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

func CreateRunWorkerPool() {
	var allTask []*workerpool.Task
	for i := 1; i <= 100; i++ {
		task := workerpool.NewTask(func(data interface{}) error {
			taskID := data.(int)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Task %d processed\n", taskID)
			return nil
		}, i)
		allTask = append(allTask, task)
	}

	pool := workerpool.NewPool(allTask, 5)
	go func() {
		for {
			taskID := rand.Intn(100) + 20

			if taskID%7 == 0 {
				pool.Stop()
			}

			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
			task := workerpool.NewTask(func(data interface{}) error {
				taskID := data.(int)
				time.Sleep(100 * time.Millisecond)
				fmt.Printf("Task %d processed\n", taskID)
				return nil
			}, taskID)
			pool.AddTask(task)
		}
	}()
	pool.RunBackground()
}
