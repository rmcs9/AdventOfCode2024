package main 


import (
    "aoc2024/data"
    "fmt"
)


func main() {
    input := data.Get("data/day15.txt")
    
    onMoves := false
    wh1 := make([][]byte, 0)
    wh2 := make([][]byte, 0)
    moves := ""
    for _, line := range input {
        if line == "" {
            onMoves = true
            continue
        }

        if !onMoves {
            mline1 := make([]byte, 0)  
            mline2 := make([]byte, 0)  
            for _, c := range line {
                mline1 = append(mline1, byte(c))
                mline2 = append(mline2, byte(c))
            }
            wh1 = append(wh1, mline1)
            wh2 = append(wh2, mline2)
        } else {
            moves += line
        }
    }

    fmt.Println("PART 1:", part1(wh1, moves)) 
    fmt.Println("PART 2:", part2(wh2, moves))
}

type vec struct {
    dx int 
    dy int
}

var vecs = map[byte]vec {
    '^':{ 0, -1 }, 
    'v':{ 0,  1 }, 
    '>':{ 1,  0 }, 
    '<':{-1,  0 },
}

func part1(wh [][]byte, moves string) int {
    var bot vec 
    for y := range wh {
        for x := range wh[y] {
            if wh[y][x] == '@' {
                bot = vec{ x, y }
                break
            }
        }
    }

    for _, m := range moves {
        dir := vecs[byte(m)]       
        if check(vec{ bot.dx + dir.dx, bot.dy + dir.dy }, dir, wh) {
            c := vec{ bot.dx, bot.dy } 
            for wh[c.dy][c.dx] != '.' {
                c.dx += dir.dx 
                c.dy += dir.dy
            }

            stop := true 
            for stop {
                if wh[c.dy - dir.dy][c.dx - dir.dx] == '@' {
                    stop = false
                }
                
                wh[c.dy][c.dx] = wh[c.dy - dir.dy][c.dx - dir.dx]
                c = vec{ c.dx - dir.dx, c.dy - dir.dy }
            }
            wh[bot.dy][bot.dx] = '.'
            bot = vec { bot.dx + dir.dx, bot.dy + dir.dy }
        }
    }

    total := 0
    for y := range wh {
        for x := range wh[y] {
            if wh[y][x] == 'O' {
                total += (y * 100) + x
            }
        }
    }
    return total
}

func check(qcoord vec, dir vec, wh [][]byte) bool {
    switch wh[qcoord.dy][qcoord.dx] {
    case '.': return true
    case '#': return false
    case 'O': return check(vec{ qcoord.dx + dir.dx, qcoord.dy + dir.dy }, dir, wh)
    }
    return false
}


func part2(smallwh [][]byte, moves string) int {
    wh := make([][]byte, 0)
    for y := range smallwh {
        wh = append(wh, make([]byte, 0))
        for x := range smallwh[y] {
            if smallwh[y][x] == '#' {
                wh[y] = append(wh[y], []byte{'#','#'}...)
            } else if smallwh[y][x] == '.' {
                wh[y] = append(wh[y], []byte{'.','.'}...)
            } else if smallwh[y][x] == 'O' {
                wh[y] = append(wh[y], '[')
                wh[y] = append(wh[y], ']')
            } else if smallwh[y][x] == '@' {
                wh[y] = append(wh[y], '@') 
                wh[y] = append(wh[y], '.') 
            }
        }
    }

    var bot vec 
    for y := range wh {
        for x := range wh[y] {
            if wh[y][x] == '@' {
                bot = vec{ x, y }
                break
            }
        }
    }

    for _, m := range moves {
        dir := vecs[byte(m)]
        if check2(vec{ bot.dx + dir.dx, bot.dy + dir.dy }, dir, wh) {
            if dir.dy == 0 {
                c := vec{ bot.dx, bot.dy } 
                for wh[c.dy][c.dx] != '.' {
                    c.dx += dir.dx 
                    c.dy += dir.dy
                }

                stop := true 
                for stop {
                    if wh[c.dy - dir.dy][c.dx - dir.dx] == '@' {
                        stop = false
                    }
                    
                    wh[c.dy][c.dx] = wh[c.dy - dir.dy][c.dx - dir.dx]
                    c = vec{ c.dx - dir.dx, c.dy - dir.dy }
                }
                wh[bot.dy][bot.dx] = '.'
                bot = vec { bot.dx + dir.dx, bot.dy + dir.dy }
            } else {
                moveUnD(bot, dir, wh)
                bot = vec{ bot.dx + dir.dx, bot.dy + dir.dy }
            }
        }
    }

    total := 0
    for y := range wh {
        for x := range wh[y] {
            if wh[y][x] == '[' {
                total += (y * 100) + x
            }
        }
    }
    return total
}

func check2(qcoord vec, dir vec, wh[][]byte) bool {
    switch wh[qcoord.dy][qcoord.dx] {
    case '.': return true
    case '#': return false
    case '[': 
        this := check2(vec{qcoord.dx + dir.dx, qcoord.dy + dir.dy}, dir, wh)
        if dir.dy != 0 {
            right := vec{ qcoord.dx + 1 + dir.dx, qcoord.dy + dir.dy } 
            return this && check2(right, dir, wh)
        }
        return this
    case ']': 
        this := check2(vec{qcoord.dx + dir.dx, qcoord.dy + dir.dy}, dir, wh)
        if dir.dy != 0 {
            left := vec{ qcoord.dx - 1 + dir.dx, qcoord.dy + dir.dy } 
            return this && check2(left, dir, wh)
        }
        return this
    }
    return false
}

//trust the plan o7
func moveUnD(coord vec, dir vec, wh [][]byte) {
    curr := vec{ coord.dx, coord.dy } 

    temp := wh[curr.dy][curr.dx]
    var fill byte = '.'

    for {
        wh[curr.dy][curr.dx] = fill 
        if temp == '.' {
            break
        } else if curr != coord && temp == '[' {
            moveUnD(vec{ curr.dx + 1, curr.dy }, dir, wh)
        } else if curr != coord && temp == ']'{
            moveUnD(vec{ curr.dx - 1, curr.dy }, dir, wh)
        }
        fill = temp
        curr = vec{ curr.dx + dir.dx, curr.dy + dir.dy }
        temp = wh[curr.dy][curr.dx]
    }
}
