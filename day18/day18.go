package main 


import (
    "aoc2024/data"
    "fmt"
    "strings"
    "strconv"
    "container/heap"
)

var memMap [][]byte
var maxX, maxY int
var falls []coord 

type coord struct {
    x int 
    y int
}

func main() {
    input := data.Get("data/day18.txt")

    for _, line := range input {
        fall := func() coord {
            xy := strings.Split(line, ",")
            var c coord 
            c.x, _ = strconv.Atoi(xy[0])
            c.y, _ = strconv.Atoi(xy[1])
            return c
        }()
        falls = append(falls, fall)
    }

    maxX, maxY = 71, 71

    for y := range maxY {
        memMap = append(memMap, make([]byte, maxX))
        for x := range maxX {
            memMap[y][x] = '.'
        }
    }

    fmt.Println("PART 1:", part1(1024)) 
    fmt.Println("PART 2:", part2())
}

func part1(bytes int) int {
    for i := range bytes {
        c := falls[i]

        memMap[c.y][c.x] = '#'
    }

    pq := make(pqueue, 0)
    seen := make(map[coord]bool)
    heap.Init(&pq)

    heap.Push(&pq, &route{coord{0,0}, 0})

    ans := -1
    for pq.Len() > 0 {
        p := heap.Pop(&pq).(*route)

        seen[p.pos] = true
        if p.pos.x == maxX - 1 && p.pos.y == maxY - 1 {
            ans = p.score
            break
        }
        
        for _, dir := range dirs {
            c := coord{p.pos.x + dir.x, p.pos.y + dir.y}

            if _, key := seen[c]; !key && inBounds(c.x, c.y) && memMap[c.y][c.x] != '#' {
                seen[c] = true
                heap.Push(&pq, &route{c, p.score + 1})       
            }
        }
    }

    for key := range seen {
        memMap[key.y][key.x] = 'O'
    }
    return ans
}

//shameless brute force :P
func part2() string {
    for i := 1025; i < len(falls); i++ {
        for y := range maxY {
            for x := range maxX {
                memMap[y][x] = '.'
            }
        }
        ans := part1(i)
        if ans == -1 {
            f := falls[i - 1]
            return fmt.Sprint(f.x, ",", f.y)
        }
    }
    return "PATH IS CLEAR"
}

var dirs = []coord {
    {0, -1}, 
    {0,  1}, 
    {-1, 0}, 
    { 1, 0},
}

type route struct {
    pos coord 
    score int
}

type pqueue []*route 

func (pq pqueue) Len() int { return len(pq) }

func (pq pqueue) Less(i, j int) bool {
    return pq[i].score < pq[j].score 
}

func (pq pqueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *pqueue) Push(x any) {
    item := x.(*route)
    *pq = append(*pq, item)
}

func (pq *pqueue) Pop() any {
    old := *pq 
    n := len(old) 
    x := old[n-1]
    old[n-1] = nil
    *pq = old[0 : n-1]
    return x
}

var inBounds = func(x, y int) bool {
    return (x >= 0 && x < maxX) && (y >= 0 && y < maxY)
}
