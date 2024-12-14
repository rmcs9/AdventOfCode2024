package main 


import (
	"aoc2024/data"
	"fmt"
	"strings"
	"strconv"
	"slices"
)

var maxX, maxY = 101, 103

func main() {
	input := data.Get("data/day14.txt")

	var bots1 []robot
	var bots2 []robot
	for _, line := range input {
		pnv := strings.Split(line, " ")
		p := pnv[0] 
		v := pnv[1]
		
		var bot1, bot2 robot 
		bot1.pos = extractCoord(p)
		bot1.vec = extractCoord(v)
		bot2.pos = extractCoord(p)
		bot2.vec = extractCoord(v)
		bots1 = append(bots1, bot1)
		bots2 = append(bots2, bot2)
	}

	fmt.Println("PART 1:", part1(bots1)) 
	fmt.Println("PART 2:", part2(bots2))
}

func part1(bots []robot) int {
	for range 100 {
		for _, bot := range bots {
			bot.pos.x += bot.vec.x 
			bot.pos.y += bot.vec.y

			if bot.pos.x > maxX - 1 {
				bot.pos.x -= maxX
			} else if bot.pos.x < 0 {
				bot.pos.x += maxX
			}
			if bot.pos.y > maxY - 1 {
				bot.pos.y -= maxY
			} else if bot.pos.y < 0 {
				bot.pos.y += maxY
			}
		}
	}

	midX := maxX / 2 
	midY := maxY / 2 

	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, bot := range bots {
		if bot.pos.x < midX && bot.pos.y < midY { q1++ }
		if bot.pos.x > midX && bot.pos.y < midY { q2++ }
		if bot.pos.x < midX && bot.pos.y > midY { q3++ }
		if bot.pos.x > midX && bot.pos.y > midY { q4++ }
	}
	return q1 * q2 * q3 * q4
}

func part2(bots []robot) int {
	tree := make([][]int, maxY) 
	for y := range maxY {
		tree[y] = make([]int, maxX)
	}

	for i := 1; ; i++ {
		for _, bot := range bots {
			bot.pos.x += bot.vec.x 
			bot.pos.y += bot.vec.y

			if bot.pos.x > maxX - 1 {
				bot.pos.x -= maxX
			} else if bot.pos.x < 0 {
				bot.pos.x += maxX
			}
			if bot.pos.y > maxY - 1 {
				bot.pos.y -= maxY
			} else if bot.pos.y < 0 {
				bot.pos.y += maxY
			}
			tree[bot.pos.y][bot.pos.x] = 1
		}

		// LMAOOOOOOOO IT WORKSSSSSSSSSSSSSSSSSSSSSSSSSS
		for y := range tree {
			for x := 0; x < len(tree[y]); x++ {
				if tree[y][x] == 1{
					found := false
					if x + 10 < len(tree[y]) {
						s := tree[y][x:x+10]
						found = slices.Min(s) == 1
					}
					if found {
						for _, line := range tree {
							for _, val := range line {
								fmt.Print(val)
							}
							fmt.Println()
						}
						return i
					}
				}
			}
		}

		tree = make([][]int, maxY) 
		for y := range maxY {
			tree[y] = make([]int, maxX)
		}
	}
}

type robot struct {
	pos *coord 
	vec *coord
}

type coord struct {
	x int 
	y int 
}

var extractCoord = func(str string) *coord {
	comma := strings.Index(str, ",")

	x := str[strings.Index(str, "=") + 1:comma] 
	str = str[comma + 1:]
	y := str

	ret := new(coord)
	ret.x, _ = strconv.Atoi(x)
	ret.y, _ = strconv.Atoi(y)
	return ret
}

var abs = func (i int) int {
	if i < 0 { i = -i } 
	return i 
}
