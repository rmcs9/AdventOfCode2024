package main 


import (
    "aoc2024/data"
    "fmt"
)

var input []string
var scores map[coord]int
var start coord
var maxX, maxY int

func main() {
    input = data.Get("data/day20.txt")

    for y := range input {
        for x := range input[y] {
            if input[y][x] == 'S' {
                start = coord{ x, y }
            }
        }
    }

    maxX, maxY = len(input[0]), len(input)

    fmt.Println("PART 1:", part1()) 
    fmt.Println("PART 2:", part2())
}

func part1() int {
    scores = make(map[coord]int)
    Q := make([]coord, 0)
    Q = append(Q, start)
    scores[start] = 0
    for len(Q) > 0 {
        p := Q[0] 
        Q = Q[1:]

        for _, dir := range dirs {
            c := coord{ p.x + dir.x, p.y + dir.y }

            if _, key := scores[c]; (input[c.y][c.x] == 'E' || input[c.y][c.x] == '.') && !key {
                scores[c] = scores[p] + 1
                Q = append(Q, c)
            }
        }
    }

    cheats := 0
    for p1, p1s := range scores {
        for _, dir := range dirs {
            p2 := coord{p1.x + dir.x, p1.y + dir.y} 
            if input[p2.y][p2.x] == '#' {
                p2 := coord{p2.x + dir.x, p2.y + dir.y} 
                if p2s, key := scores[p2]; key {
                    if p2s > p1s {
                        if (p2s - p1s) - 2 >= 100 {
                            cheats++
                        }
                    }
                }
            }
        }
    }
    return cheats
}

// need to id all the points that are <= 20 ps away from u 
// and then figure out if each of those points save you at least 100 ps
func part2() int {
    cheats := 0
    for p1, p1s := range scores {
        for p2, p2s := range scores {
            if p1 == p2 || p1s > p2s {
                continue
            }
            dx := abs(p1.x - p2.x)
            dy := abs(p1.y - p2.y)
            ds := (p2s - p1s) - dx - dy
            if dx + dy <= 20 && ds >= 100 {
                cheats++
            }
        }
    }
    return cheats
}

type coord struct {
    x int 
    y int
}

var dirs = []coord {
    {0, -1}, 
    {0,  1}, 
    {-1, 0}, 
    { 1, 0},
}

var inBounds = func(x, y int) bool {
    return (x >= 0 && x < maxX) && (y >= 0 && y < maxY)
}

var abs = func (i int) int {
    if i < 0 { i = -i } 
    return i 
}
