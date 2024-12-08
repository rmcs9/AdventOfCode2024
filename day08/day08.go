package main

import ( 
	"aoc2024/data"
	"fmt"
)

var input []string
var sigs map[rune][]coord
var locs map[coord]bool
var maxX, maxY int

func main() {
	input = data.Get("data/day08.txt")
	sigs = make(map[rune][]coord)

	maxX = len(input[0])
	maxY = len(input)

	for y, line := range input {
		for x, sym := range line {
			if sym != '.' {
				_, key := sigs[sym]
				if !key {
					sigs[sym] = make([]coord, 0)
				}
				sigs[sym] = append(sigs[sym], coord{ x, y })
			}
		}
	}

	fmt.Println("PART 1:", part1())
	fmt.Println("PART 2:", part2())
}

func part1() int {
	locs = make(map[coord]bool) 
	for _, coords := range sigs {
		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				dX := abs(coords[i].X - coords[j].X)
				dY := abs(coords[i].Y - coords[j].Y)

				iX, iY := coords[i].X, coords[i].Y
				jX, jY := coords[j].X, coords[j].Y

				if iX > jX {
					iX += dX 
					jX -= dX
				} else {
					iX -= dX 
					jX += dX
				}

				if iY > jY {
					iY += dY 
					jY -= dY
				} else {
					iY -= dY 
					jY += dY
				}

				if inBounds(iX, iY) { locs[coord { iX, iY }] = true }
				if inBounds(jX, jY) { locs[coord { jX, jY }] = true }
			}
		}
	}
	return len(locs)
}

func part2() int {
	locs = make(map[coord]bool)
	for _, coords := range sigs {
		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				dX := abs(coords[i].X - coords[j].X)
				dY := abs(coords[i].Y - coords[j].Y)

				iX, iY := coords[i].X, coords[i].Y
				jX, jY := coords[j].X, coords[j].Y

				determineDirection := func (x, y int) (bool, bool) {
					return x > y, !(x > y)
				}

				iXV, jXV := determineDirection(iX, jX)
				iYV, jYV := determineDirection(iY, jY)

				extend := func(x, y int, xV, yV bool) {
					for inBounds(x, y) {
						locs[coord { x, y }] = true
						if xV {
							x += dX
						} else {
							x -= dX
						}

						if yV {
							y += dY
						} else {
							y -= dY
						}
					}				
				}

				extend(iX, iY, iXV, iYV)
				extend(jX, jY, jXV, jYV)
			}
		}
	}
	return len(locs)
}

var inBounds = func(x, y int) bool {
	return (x >= 0 && x < maxX) && (y >= 0 && y < maxY)
}

var abs = func (i int) int {
	if i < 0 { i = -i } 
	return i 
}

type coord struct {
	X int
	Y int
}
