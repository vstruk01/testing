package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/vstruk01/testing/popcount"
	"github.com/vstruk01/testing/workerpool"
)

func main() {
	// x, y := 5, 3
	// fmt.Println(x &^ y)

	fmt.Println(popcount.PopCountCycle(23423935))
	fmt.Println(popcount.PopCountHard(23423935))
	fmt.Println(popcount.PopCountEachBit(23423935))

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
