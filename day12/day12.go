package main 


import (
    "aoc2024/data"
    "fmt"
)

var input []string
var maxX, maxY int

func main() {
    input = data.Get("data/day12.txt")

    maxY = len(input)
    maxX = len(input[0])

    fmt.Println("PART 1:", part1()) 
    fmt.Println("PART 2:", part2())
}

func part1() int {
    foundPoints := make(map[coord]bool)
    total := 0
    for y := range input {
        for x := range input[y] {
            _, in := foundPoints[coord{ x, y }]
            if !in {
                total += func() int {
                    subFound := make(map[coord]bool)
                    Q := make([]coord, 0)
                    Q = append(Q, coord { x, y })
                    sym := input[y][x]
                    area, perim := 0, 0
                    for len(Q) > 0 {
                        np := Q[0]
                        Q = Q[1:]
                        if _, f := subFound[np]; !f {
                            area++ 
                            newPoints := 0
                            for _, dir := range dirs {
                                nx, ny := np.x + dir.x, np.y + dir.y
                                if inBounds(nx, ny) && input[ny][nx] == sym {
                                    newPoints++
                                    Q = append(Q, coord{ nx, ny })
                                }
                            }

                            perim += 4 - newPoints
                        }
                        subFound[np] = true
                        foundPoints[np] = true
                    }
                    return area * perim
                }()
            }
        }
    }
    return total
}

func part2() int {
    foundPoints := make(map[coord]bool)
    total := 0
    for y := range input {
        for x := range input[y] {
            _, in := foundPoints[coord{ x, y }]
            if !in {
                total += func() int {
                    subFound := make(map[coord]bool)
                    Q := make([]coord, 0)
                    Q = append(Q, coord { x, y })
                    sym := string(input[y][x])
                    area, sides := 0, 0
                    for len(Q) > 0 {
                        np := Q[0]
                        Q = Q[1:]
                        if _, f := subFound[np]; !f {
                            area++ 
                            for _, dir := range dirs {
                                nx, ny := np.x + dir.x, np.y + dir.y
                                if inBounds(nx, ny) && string(input[ny][nx]) == sym {
                                    Q = append(Q, coord{ nx, ny })
                                }
                            }

                            // the idea here is at every point check the orthogonally adjacent points 
                            // as well as the point that is diagonal to ur point.
                            //
                            // u have an exterior corner if the orthogonal points are not part of the region, 
                            //
                            // u have an interior corner if the orthogonal points are part of the region, 
                            // but the diagonal point is not.
                            //
                            // IT IS IMPORTANT TO ONLY COUNT A CORNER AS ONE SIDE. I WAS COUNTING INTERIOR CORNERS 
                            // AS 2 SIDES FOR 30 MINUTES TRYING TO FIGURE OUT WHAT WAS WRONG
                            for _, dir := range cornerdirs {
                                l, r, d := "_", "_", "_"
                                if inBounds(np.x + dir[0].x, np.y + dir[0].y) {
                                    l = string(input[np.y + dir[0].y][np.x + dir[0].x])
                                }
                                if inBounds(np.x + dir[1].x, np.y + dir[1].y) {
                                    r = string(input[np.y + dir[1].y][np.x + dir[1].x])
                                }
                                if inBounds(np.x + dir[2].x, np.y + dir[2].y) {
                                    d = string(input[np.y + dir[2].y][np.x + dir[2].x])
                                }
                                if (l != sym && r != sym) || (l == sym && r == sym && d != sym) {
                                    sides++
                                }
                            }
                        }
                        subFound[np] = true
                        foundPoints[np] = true
                    }
                    return area * sides
                }()
            }
        }
    }
    return total
}

type coord struct {
    x int 
    y int
}

var dirs = []coord {
    //north 
    {0, -1},
    //south 
    {0,  1},
    //east 
    {1,  0},
    //west 
    {-1, 0},
}

var cornerdirs = [][3]coord {
    //topright corner north, east, ne
    { {0, -1}, {1,  0}, {1, -1 } }, 
    //bottomright corner south, east, se
    { {0,  1}, {1,  0}, {1, 1} },
    //bottomleft corner south, west, sw 
    { {0,  1}, {-1, 0}, {-1, 1 } },  
    //topleft corner north, west, nw
    { {0, -1}, {-1, 0}, {-1, -1 } },
} 

var inBounds = func(x, y int) bool {
    return (x >= 0 && x < maxX) && (y >= 0 && y < maxY)
}
