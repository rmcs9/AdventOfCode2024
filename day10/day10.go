package main 

import (
	"aoc2024/data"
	"fmt"
	"strconv"
	"log"
)

var grid [][]int
var heads map[coord]bool
var maxX, maxY int
var vecs  = []vec{
	//north
	{0, -1},
	//south 
	{0,  1}, 
	//east 
	{1,  0},
	//west
	{-1, 0},
}

func main() {
	input := data.Get("data/day10.txt")
	heads = make(map[coord]bool)

	maxX, maxY = len(input[0]), len(input)
	
	grid = make([][]int, 0)
	for i := range input {
		grid = append(grid, make([]int, 0))
		for j := range input[i] {
			if input[i][j] == '0' {
				heads[coord{ j, i }] = true	
			}
			num, err := strconv.Atoi(string(input[i][j]))
			if err != nil {
				log.Fatal(err)
			}
			grid[i] = append(grid[i], num)
		}
	}

	fmt.Println("PART 1", part1())
	fmt.Println("PART 2", part2())
}

func part1() int {
	total := 0
	for head := range heads {
		mp := new(map[coord]bool)
		(*mp) = make(map[coord]bool)
		total += determineTrails(head, 0, mp)
	}
	return total
}

func determineTrails(p coord, num int, m *map[coord]bool) int {
	points := make([]coord, 0)
	for _, vec := range vecs {
		newX, newY := p.x + vec.dx, p.y + vec.dy
		if inBounds(newX, newY) && grid[newY][newX] == num + 1 {
			points = append(points, coord{ newX,  newY })
		}
	}

	if num == 8 {
		for _, point := range points {
			(*m)[point] = true
		}
	} else {
		for _, point := range points {
			determineTrails(point, num + 1, m)
		}
	}
	return len(*m)
}

func part2() int {
	total := 0
	for head := range heads {
		total += determineTrails2(head, 0)
	}
	return total
}

func determineTrails2(p coord, num int) int {
	points := make([]coord, 0)
	for _, vec := range vecs {
		newX, newY := p.x + vec.dx, p.y + vec.dy
		if inBounds(newX, newY) && grid[newY][newX] == num + 1 {
			points = append(points, coord{ newX,  newY })
		}
	}

	if num == 8 {
		return len(points)
	}

	total := 0
	for i := range points {
		total += determineTrails2(points[i], num + 1) 
	}
	return total
}

var inBounds = func(x, y int) bool {
	return (x >= 0 && x < maxX) && (y >= 0 && y < maxY)
}


type coord struct {
	x int
	y int
}

type vec struct {
	dx int 
	dy int
}
