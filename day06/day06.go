package main


import ( 
	"aoc2024/data"
	"fmt"
	"strings"
	"sync"
)

var input []string
var guardPath map[vector]bool
var guardPathDir map[vectorPoint]bool 
var knownObstacles = sync.Map{}

var startPoint vector

var vectors = []vector {
	//up 
	{0, -1},	
	//right
	{ 1, 0},
	//down 
	{0,  1},
	//left
	{-1, 0}, 
}

type vector struct {
	Dx int 
	Dy int 
}

type vectorPoint struct {
	X int 
	Y int 
	V int
}
var inBounds = func (x, y int) bool {
	return (x < len(input[0]) && x >= 0) && (y < len(input) && y >= 0)
}

var changeDir = func (dir *int) {
	if *dir == 3 {
		*dir = 0
	} else {
		*dir++
	}
}

func main() {
	input = data.Get("data/day06.txt")

	fmt.Println("PART 1: ", part1())
	// interestingly threads past 100 actually start to slow down execution
	// probably due to the size of the lists and/or thread construction time
	// TIMES: 
	// 1 thread:    3.353 sec
	// 100 threads:  .377 sec
	// 150 threads:  .389 sec 
	// 200 threads:  .392 sec
	fmt.Println("PART 2: ", part2(100))
}

func part1() int {
	//find the starting point
	for i := range input {
		j := strings.Index(input[i], "^")	
		if j != -1 {
			startPoint.Dx = j
			startPoint.Dy = i
			break
		}
	}

	x, y, d, dir := startPoint.Dx, startPoint.Dy, vectors[0], 0

	// vector struct reused here to mark coords
	guardPath = make(map[vector]bool)
	guardPathDir = make(map[vectorPoint]bool)

	for inBounds(x, y) {
		if input[y][x] == '#' {
			//take a step back
			y = y - d.Dy
			x = x - d.Dx

			//change direction
			changeDir(&dir) 
			d = vectors[dir]
		} else {
			guardPath[vector{x, y}] = true
			guardPathDir[vectorPoint{x, y, dir}] = true

			y = y + d.Dy 
			x = x + d.Dx
		}
	}

	return len(guardPath)
}

// ASYNC DAY!!!!!!!!!!!!
func part2(threads int) int {
	total := 0

	// convert the mapped points into a slice
	pointsSlice := make([]vectorPoint, 0, len(guardPathDir))
	for key := range guardPathDir {
		pointsSlice = append(pointsSlice, key)
	}
	
	//calc the subsize for each drone
	subSize := len(guardPathDir) / threads

	var lists [][]vectorPoint
	
	//divide the points into subsized lists
	for i := range threads - 1  {
		lists = append(lists, make([]vectorPoint, 0))
		lists[i] = append(lists[i], pointsSlice[:subSize]...)
		pointsSlice = pointsSlice[subSize:]
	}

	lists = append(lists, make([]vectorPoint, 0))
	lists[threads - 1] = append(lists[threads - 1], pointsSlice...)

	//prepare the channels
	var channels []chan int
	for range threads {
		channels = append(channels, make(chan int))
	}

	// execute the drones
	for i := range threads {
		go p2drone(lists[i], channels[i])
	}

	//collect values from the channels
	for i := range threads {
		total += <-channels[i]
	}

	return total
}

func p2drone(points []vectorPoint, ch chan int) {
	total := 0
	for _, point := range points {
		
		newPath := make(map[vectorPoint]int)

		x, y, d, dir := startPoint.Dx, startPoint.Dy, vectors[0], 0
		
		obsPoint := vector{point.X + vectors[point.V].Dx, point.Y + vectors[point.V].Dy}
		
		known := false
		for inBounds(x, y) {
			if input[y][x] == '#' || (y == obsPoint.Dy && x == obsPoint.Dx) {
				//take a step back
				y = y - d.Dy
				x = x - d.Dx

				//change direction
				changeDir(&dir) 
				d = vectors[dir]
			} else {
				_, key := newPath[vectorPoint{x, y, dir}]
				if key {
					// check if this obstacle already caused another loop
					_, key := knownObstacles.Load(obsPoint)
					known = key
					break
				}
				newPath[vectorPoint{x, y, dir}] = newPath[vectorPoint{x, y, dir}] + 1 
				y = y + d.Dy 
				x = x + d.Dx
			}		
		}

		if inBounds(x, y) && !known { 
			knownObstacles.Store(obsPoint, true)
			total++ 
		}
	}
	ch <- total
	close(ch)
}
