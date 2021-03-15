package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/vstruk01/testing/popcount"
	"github.com/vstruk01/testing/workerpool"
	// "github.com/vstruk01/testing/workerpool"
)

// "github.com/vstruk01/testing/workerpool"

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
	pool.RunBackground()

	pool.AddTask(workerpool.NewTask(func(data interface{}) error {
		taskID := data.(int)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Task %d processed\n", taskID)
		return nil
	}, 244))
	pool.Stop()

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

// const (
// 	width, height = 600, 320            // canvas size in pixels
// 	cells         = 100                 // number of grid cells
// 	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
// 	xyscale       = width / 2 / xyrange // pixels per x or y unit
// 	zscale        = height * 0.4        // pixels per z unit
// 	angle         = math.Pi / 6         // angle of x, y axes (=30°)
// )

// var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// func main() {
// 	http.Handle("/getCorner", GetCorner)

// 	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
// 		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
// 		"width='%d' height='%d'>", width, height)
// 	for i := 0; i < cells; i++ {
// 		for j := 0; j < cells; j++ {
// 			ax, ay := corner(i+1, j)
// 			bx, by := corner(i, j)
// 			cx, cy := corner(i, j+1)
// 			dx, dy := corner(i+1, j+1)
// 			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
// 				ax, ay, bx, by, cx, cy, dx, dy)
// 		}
// 	}
// 	fmt.Println("</svg>")
// }

// func GetCorner(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("ContentType", "image/svg+xml")
// }

// func corner(i, j int) (float64, float64) {
// 	// Find point (x,y) at corner of cell (i,j).
// 	x := xyrange * (float64(i)/cells - 0.5)
// 	y := xyrange * (float64(j)/cells - 0.5)

// 	// Compute surface height z.
// 	z := f(x, y)

// 	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
// 	sx := width/2 + (x-y)*cos30*xyscale
// 	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
// 	return sx, sy
// }

// func f(x, y float64) float64 {
// 	r := math.Hypot(x, y) // distance from (0,0)
// 	return math.Sin(r) / r
// }
