package main 


import (
    "aoc2024/data"
    "fmt"
    "container/heap"
)

var input []string 

var start coord

func main() {
    input = data.Get("data/day16.txt")

    for y := range input {
        for x := range input[y] {
            if input[y][x] == 'S' {
                start.x = x 
                start.y = y
            }
        }
    }

    fmt.Println("PART 1:", part1()) 
    fmt.Println("PART 2:", part2())
}

func part1() int {
    pq := make(pqueue, 0)
    heap.Init(&pq)
    seen := make(map[state]int)
    var end route

    heap.Push(&pq, &route{start, 0, coord{1, 0}})

    for pq.Len() > 0 {

        p := heap.Pop(&pq).(*route)
        if input[(*p).pos.y][(*p).pos.x] == 'E' {
            end = *p
            break
        }

        for _, dir := range dirs {
            c := coord{p.pos.x + dir.x, p.pos.y + dir.y} 
            score := (*p).score
            if dir == p.dir {
                score += 1
            } else { 
                score += 1001 
            }
            if val, key := seen[state{c, dir}]; (!key || score < val) && input[c.y][c.x] != '#' {
                heap.Push(&pq, &route{c, score, dir})
            }
        }
		seen[state{p.pos, p.dir}] = p.score
    }
    return end.score
}

func part2() int {
    pq := make(pqueue, 0)
    heap.Init(&pq)
    seen := make(map[state]int)
    // map from state{ which is a point and a direction } -> parent { lowest score, and a collection of parent points to this point }
    // creates a kind of closed tree structure with start and end at the ends(2 roots) and many branching points in the middle
    // caching with direction is important here, as it will allow for multiple paths of the same length to be logged
    parents := make(map[state]*parent)
    var end state

    heap.Push(&pq, &route{start, 0, coord{1, 0}})
    startState := state{start, coord{1,0}}
    parents[startState] = &parent { 0, make([]state, 0) }

    for pq.Len() > 0 {

        p := heap.Pop(&pq).(*route)
        if input[(*p).pos.y][(*p).pos.x] == 'E' {
            end = state{(*p).pos, (*p).dir}
            break
        }

        for _, dir := range dirs {
            c := coord{p.pos.x + dir.x, p.pos.y + dir.y}
            score := p.score
            if dir == p.dir {
                score += 1
            } else {
                score += 1001
            }

            //check to see if the node you are travelling to has been visited 
            if val, key := parents[state{c, dir}]; key {
                //if it has been visited, check if the current points score beats out the previous visitor
                if score < val.score {
                    // if it does, reset the score and set the new parents to be the current point and direction
                    val.score = score 
                    val.parents = []state{{(*p).pos, (*p).dir}}
                } else if score == val.score {
                    //if the scores match, add the current point to the points parents
                    val.parents = append(val.parents, state{(*p).pos, (*p).dir})
                }
            } else {
                // if the new point has not been visited, create a new parent struct for it
                parents[state{c, dir}] = &parent{ score, []state{ {(*p).pos, (*p).dir}}}
            }

            // standard p1 dijkstras stuff
            if val, key := seen[state{c, dir}]; (!key || score < val ) && input[c.y][c.x] != '#' {
                heap.Push(&pq, &route{c, score, dir})
		        seen[state{c, dir}] = p.score
            }
        }
    }
    
    // after reaching the end point and creating the 2 rooted tree,
    // walk the entire tree from the end point and collect every unique coord 
    pathSeen := make(map[coord]bool)
    q := []state{ end }
    var z state
    for len(q) > 0 {
        c := q[0] 
        q = q[1:]
        if c != z {
            pathSeen[c.pos] = true
            q = append(q, parents[c].parents...)
        }
    } 
    return len(pathSeen)
}

type coord struct {
    x int 
    y int 
}

type route struct {
    pos coord 
    score int
    dir coord
}

type state struct {
    pos coord 
    dir coord
}

var dirs = []coord {
    {0, -1}, 
    {0,  1}, 
    {-1, 0}, 
    { 1, 0},
}

type parent struct {
    score int 
    parents []state
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
