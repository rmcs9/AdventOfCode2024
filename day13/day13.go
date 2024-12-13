package main 


import (
	"aoc2024/data"
	"fmt"
	"strings"
	"strconv"
)

var input []string
var claws []machine 

func main() {
	input = data.Get("data/day13.txt")
	claws = make([]machine, 0)

	for i := 0; i < len(input); i += 4 {
		Abutton := input[i] 
		Bbutton := input[i + 1] 
		Prize := input[i + 2]
		ax, ay := extractButton(Abutton)
		bx, by := extractButton(Bbutton)
		px, py := extractPrize(Prize)

		m := makeMachine(ax, ay, bx, by, px, py)
		claws = append(claws, m)
	}

	fmt.Println("PART 1:", part1()) 
	fmt.Println("PART 2:", part2())
}

// solve the equations: 
//
// ax + bx = px 
// ay + by = py 
//
// for both parts 
func part1() int {
	tokens := 0
	for _, claw := range claws {
		// ay(ax + bx) = ay(px)
		nbX := claw.b.x * claw.a.y
		npX := claw.prize.x * claw.a.y 
		// ax(ay + by) = ax(py)
		nbY := claw.b.y * claw.a.x 
		npY := claw.prize.y * claw.a.x

		// ay(ax + bx) - ax(ay + by) = ay(px) - ax(py)
		// cancels out the a values
		B := abs(nbX - nbY) 
		P := abs(npX - npY)
		
		// ensure that (ay(px) - ax(py)) / aybx - axby 
		// is evenly divisible 
		if P % B == 0 {
			//solve for b
			b := P / B
			// with the newfound value for b, plug in to solve for a
			if (claw.prize.x - (b * claw.b.x)) % claw.a.x == 0 {
				a := ((claw.prize.x - (b * claw.b.x)) / claw.a.x) 
				// ensure that a and b are within range
				if a <= 100 && b <= 100 {
					tokens += (a * 3) + b
				}
			}
			
		}
	}
	return tokens
}

// same as part 1 but add 10 bil
func part2() int {	
	tokens := 0
	for _, claw := range claws {
		claw.prize.x += 10000000000000
		claw.prize.y += 10000000000000
		nbX := claw.b.x * claw.a.y
		nbY := claw.b.y * claw.a.x 
		npX := claw.prize.x * claw.a.y 
		npY := claw.prize.y * claw.a.x

		B := abs(nbX - nbY) 
		P := abs(npX - npY)

		if P % B == 0 {
			b := P / B
			if (claw.prize.x - (b * claw.b.x)) % claw.a.x == 0 {
				a := ((claw.prize.x - (b * claw.b.x)) / claw.a.x) 
				tokens += (a * 3) + b
			}
		}
	}
	return tokens
}

var abs = func (i int) int {
	if i < 0 { i = -i } 
	return i 
}

type machine struct {
	a button
	b button 
	prize button
}

type button struct {
	x int 
	y int 
}

func makeMachine(ax, ay, bx, by, px, py int) machine {
	var m machine
	m.a.x = ax 
	m.a.y = ay 
	m.b.x = bx 
	m.b.y = by 
	m.prize.x = px 
	m.prize.y = py
	return m
}

func extractButton(str string) (x, y int) {
	comma := strings.Index(str, ",")
	xStr := str[strings.Index(str, "+") + 1: comma]
	str = str[comma + 1:]
	yStr := str[strings.Index(str, "+") + 1:]
	x, _ = strconv.Atoi(xStr)
	y, _ = strconv.Atoi(yStr)
	return
}

func extractPrize(str string) (x, y int) {
	comma := strings.Index(str, ",")
	xStr := str[strings.Index(str, "=") + 1: comma] 
	str = str[comma + 1:]
	yStr := str[strings.Index(str, "=") + 1:]
	x, _ = strconv.Atoi(xStr) 
	y, _ = strconv.Atoi(yStr)
	return 
}
