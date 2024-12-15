package main

import (
    "aoc2024/data"
    "fmt"
)

var input []string

var vectors = []vector {
    //left
    {-1, 0}, 
    //right
    { 1, 0},
    //top 
    {0, -1},
    //bottom 
    {0,  1},

    //topright 
    {1, -1},
    //topleft
    {-1,-1},
    //bottomright
    { 1, 1},
    //bottom left 
    {-1, 1},
}

type vector struct {
    Dx int 
    Dy int 
}

func main() {
    input = data.Get("data/day04.txt")

    fmt.Println("PART 1: ", part1())
    fmt.Println("PART 2: ", part2())
}

func part1() int {
    sum := 0
    for y := range input {
        for x := range input[y] {
            if input[y][x] == 'X' {
                sum += findXMAS1(x,y)	
            }
        }
    }
    return sum
}

func findXMAS1(x, y int) int {
    left := x >= 3 
    right := x <= len(input[y]) - 4
    top := y >= 3 
    bottom := y <= len(input) - 4

    topright := top && right
    topleft := top && left
    bottomright := bottom && right
    bottomleft := bottom && left

    var vecs []vector
    var xmas = [3]byte { 'M', 'A', 'S' }

    if left        { vecs = append(vecs, vectors[0]) } 
    if right       { vecs = append(vecs, vectors[1]) } 
    if top         { vecs = append(vecs, vectors[2]) } 
    if bottom      { vecs = append(vecs, vectors[3]) } 
    if topright    { vecs = append(vecs, vectors[4]) } 
    if topleft     { vecs = append(vecs, vectors[5]) } 
    if bottomright { vecs = append(vecs, vectors[6]) } 
    if bottomleft  { vecs = append(vecs, vectors[7]) } 


    return func() int {
        validDirections := 0 
        for _, v := range vecs {

            hold := true 
            nx := x + v.Dx 
            ny := y + v.Dy 
            for i := range 3 {
                if input[ny][nx] != xmas[i] {
                    hold = false
                    break
                }
                nx += v.Dx 
                ny += v.Dy
            }
            if hold { validDirections ++ }
        }
        return validDirections
    }()
}


func part2() int {
    sum := 0
    for y := range input {
        for x := range input[y] {
            if input[y][x] == 'A' {
                sum += findXMAS2(x,y)	
            }
        }
    }
    return sum
}

func findXMAS2(x, y int) int {
    left := x >= 1 
    right := x <= len(input[y]) - 2
    top := y >= 1 
    bottom := y <= len(input) - 2

    tr, tr_vec := top && right, vectors[4] 
    tl, tl_vec := top && left, vectors[5]
    br, br_vec := bottom && right, vectors[6]
    bl, bl_vec := bottom && left, vectors[7]

    //if any corner is unreachable an X-MAS is not possible
    if !tr || !tl || !br || ! bl {
        return 0
    }

    tr_val := input[y + tr_vec.Dy][x + tr_vec.Dx]
    tl_val := input[y + tl_vec.Dy][x + tl_vec.Dx]
    br_val := input[y + br_vec.Dy][x + br_vec.Dx]
    bl_val := input[y + bl_vec.Dy][x + bl_vec.Dx]

    //opposite diagonal corners must be exclusive to eachother
    diag1 := (tr_val == 'M' && bl_val == 'S') || (tr_val == 'S' && bl_val == 'M')
    diag2 := (tl_val == 'M' && br_val == 'S') || (tl_val == 'S' && br_val == 'M')

    if diag1 && diag2 {
        return 1 
    }
    return 0
}
