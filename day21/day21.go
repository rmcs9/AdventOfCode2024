package main

import (
    "aoc2024/data"
    "fmt"
    "log"
    "strconv"
    "strings"
    "math"
)

var input []string
var numPaths map[string]string
var dPaths map[string]string


type coord struct {
    x, y int
}

func main() {
    input = data.Get("data/day21.txt")

    // determine the number pad weights
    syms := map[coord]byte { 
        {2, 3}:'A',
        {1, 3}:'0', 
        {0, 2}:'1', 
        {1, 2}:'2', 
        {2, 2}:'3', 
        {0, 1}:'4', 
        {1, 1}:'5', 
        {2, 1}:'6', 
        {0, 0}:'7', 
        {1, 0}:'8', 
        {2, 0}:'9',
    }
    numPaths = make(map[string]string)

    for from, fsym := range syms {
        for to, tsym := range syms {
            if from != to {
                tx, ty := to.x, to.y  
                fx, fy := from.x, from.y

                dx := tx - fx 
                dy := ty - fy

                vpath := ""
                if dy < 0 { 
                    for range -dy {
                        vpath += "^"
                    }
                } else if dy > 0 {
                    for range dy {
                        vpath += "v"
                    }
                }
                hpath := ""
                if dx < 0 {
                    for range -dx {
                        hpath += "<" 
                    }
                } else if dx > 0 {
                    for range dx {
                        hpath += ">"
                    }
                }

                if _, key := syms[coord{fx,ty}]; dx > 0 && key {
                    numPaths[string(fsym) + string(tsym)] = vpath + hpath
                } else if _, key := syms[coord{tx, fy}]; key {
                    numPaths[string(fsym) + string(tsym)] = hpath + vpath
                } else if _, key := syms[coord{fx,ty}]; key {
                    numPaths[string(fsym) + string(tsym)] = vpath + hpath
                }
            }
        }
    }

    //determine the directional pad weights 

    dPaths = make(map[string]string)
    dsyms := map[coord]byte {
        {1, 0}:'^', 
        {2, 0}:'A', 
        {0, 1}:'<',
        {1, 1}:'v', 
        {2, 1}:'>',
    }

    for from, fsym := range dsyms {
        for to, tsym := range dsyms {
            if from != to {
                tx, ty := to.x, to.y  
                fx, fy := from.x, from.y

                dx := tx - fx 
                dy := ty - fy

                vpath := ""
                if dy < 0 { 
                    for range -dy {
                        vpath += "^"
                    }
                } else if dy > 0 {
                    for range dy {
                        vpath += "v"
                    }
                }
                hpath := ""
                if dx < 0 {
                    for range -dx {
                        hpath += "<" 
                    }
                } else if dx > 0 {
                    for range dx {
                        hpath += ">"
                    }
                }

                if _, key := syms[coord{fx,ty}]; dx > 0 && key {
                    dPaths[string(fsym) + string(tsym)] = vpath + hpath
                } else if _, key := syms[coord{tx, fy}]; key {
                    dPaths[string(fsym) + string(tsym)] = hpath + vpath
                } else if _, key := syms[coord{fx,ty}]; key {
                    dPaths[string(fsym) + string(tsym)] = vpath + hpath
                }
            }
        }
    }

    fmt.Println("PART 1:", part1()) 
    fmt.Println("PART 2:", part2())
}

var memo map[mem]int

type mem struct {
    str string 
    lvl int
}

func part1() int {
    total := 0
    for _, code := range input {
        total += func() int {
            num, err := strconv.Atoi(code[:strings.Index(code, "A")])
            if err != nil {
                log.Fatal(err)
            }
            path := operateNumPad(code, 2)
            return num * len(path)
        }()
    }
    return total
}

func operateNumPad(code string, robots int) string {
    path := ""
    from := 'A'
    for _, to := range code {
        fromto := string(from) + string(to) 
        path += numPaths[fromto] + "A"
        from = to
    }
    return operateDPad(path, robots, 1)
}

func operateDPad(path string, robots, level int) string {
    npath := "" 
    from := 'A' 
    for _, to := range path {
        if from == to { 
            npath += "A" 
            continue
        }
        fromto := string(from) + string(to)
        npath += dPaths[fromto] + "A" 
        from = to
    }
    if level == robots { return npath } 
    return operateDPad(npath, robots, level + 1)
}


// it turns out that for some reason, if the amount of robots exceeds 2 or more, 
// the ability to predetermine a single sound route between each key is no longer possible
// there will be some silly edge cases in between the robots (i assume the further u abstract out) 
// that will somehow make a different path between 2 keys faster than the one you already determined
// so, all possible paths between the keys were hardcoded, and part 2 decides which one is the cheapest and 
// selects that one. 

// really is a shame because i sunk SO MUCH time into determining the weights between keys and really thought
// i had it

// if I were to start fresh on this problem, I would determine the paths at each step instead of all at once in the beginning :(
func part2() int {
    total := 0
    for _, code := range input {
        memo = make(map[mem]int)
        total += func() int {
            num, err := strconv.Atoi(code[:strings.Index(code, "A")]) 
            if err != nil {
                log.Fatal(err)
            }

            path := operateNumPad2(code, 26, 0, numpad)
            return num * path
        }()
    }
    return total
}

func operateNumPad2(path string, robots, depth int, pad map[rune]map[rune][]string)int {

    if depth >= robots {
        return len(path)
    }

    m := mem{path, depth}
    if val := memo[m]; val != 0 {
        return val
    }

    total := 0
    from := 'A'
    for _, to := range path {
        l := math.MaxInt

        for _, o := range pad[from][to] {
            l = min(l, operateNumPad2(o, robots, depth + 1, dirpad))
        }
        total += l
        from = to
    }

    memo[m] = total
    return total
}


var numpad = map[rune]map[rune][]string{
    'A': {
        '7': {"^^^<<A", "<^^^<A"}, '8': {"^^^<A", "<^^^A"}, '9': {"^^^A"},
        '4': {"^^<<A"}, '5': {"^^<A", "<^^A"}, '6': {"^^A"},
        '1': {"^<<A"}, '2': {"^<A", "^<A"}, '3': {"^A"},
        '0': {"<A"}, 'A': {"A"},
    },
    '0': {
        '7': {"^^^<A", "^^^<A"}, '8': {"^^^A"}, '9': {"^^^>A", ">^^^A"},
        '4': {"^^<A"}, '5': {"^^A"}, '6': {"^^>A", ">^^A"},
        '1': {"^<A", "^<A"}, '2': {"^A"}, '3': {"^>A", ">^A"},
        '0': {"A"}, 'A': {">A"},
    },
    '1': {
        '7': {"^^A"}, '8': {"^^>A", ">^^A"}, '9': {"^^>>A", ">>^^A"},
        '4': {"^A"}, '5': {"^>A", ">^A"}, '6': {"^>>A", ">>^A"},
        '1': {"A"}, '2': {">A"}, '3': {">>A"},
        '0': {">vA"}, 'A': {">>vA"},
    },
    '2': {
        '7': {"^^<A", "<^^A"}, '8': {"^^A"}, '9': {"^^>A", ">^^A"},
        '4': {"^<A", "<^A"}, '5': {"^A"}, '6': {"^>A", ">^A"},
        '1': {"<A"}, '2': {"A"}, '3': {">A"},
        '0': {"vA"}, 'A': {">vA", "v>A"},
    },
    '3': {
        '7': {"^^<<A", "<<^^A"}, '8': {"^^<A", "<^^A"}, '9': {"^^A"},
        '4': {"^<<A", "<<^A"}, '5': {"^<A", "<^A"}, '6': {"^A"},
        '1': {"<<A"}, '2': {"<A"}, '3': {"A"},
        '0': {"<vA", "v<A"}, 'A': {"vA"},
    },
    '4': {
        '7': {"^A"}, '8': {"^>A", ">^A"}, '9': {"^>>A", ">>^A"},
        '4': {"A"}, '5': {">A"}, '6': {">>A"},
        '1': {"vA"}, '2': {">vA", "v>A"}, '3': {">>vA", "v>>A"},
        '0': {">vvA"}, 'A': {">>vvA"},
    },
    '5': {
        '7': {"^<A", ">^A"}, '8': {"^A"}, '9': {"^>A", "v^A"},
        '4': {"<A"}, '5': {"A"}, '6': {">A"},
        '1': {"<vA", "v<A"}, '2': {"vA"}, '3': {">vA", "v>A"},
        '0': {"vvA"}, 'A': {">vvA", "vv>A"},
    },
    '6': {
        '7': {"^<<A", "<<^A"}, '8': {"^<A", "<^A"}, '9': {"^A"},
        '4': {"<<A"}, '5': {"<A"}, '6': {"A"},
        '1': {"<<vA", "v<<A"}, '2': {"<vA", "<vA"}, '3': {"vA"},
        '0': {"<vvA", "vv<A"}, 'A': {"vvA"},
    },
    '7': {
        '7': {"A"}, '8': {">A"}, '9': {">>A"},
        '4': {"vA"}, '5': {">vA", "v>A"}, '6': {">>vA", "v>>A"},
        '1': {"vvA"}, '2': {">vvA", "vv>A"}, '3': {">>vvA", "vv>>A"},
        '0': {">vvvA"}, 'A': {">>vvvA"},
    },
    '8': {
        '7': {"<A"}, '8': {"A"}, '9': {">A"},
        '4': {"<vA", "v<A"}, '5': {"vA"}, '6': {">vA", "v>A"},
        '1': {"<vvA", "vv<A"}, '2': {"vvA"}, '3': {">vvA", "vvA"},
        '0': {"vvvA"}, 'A': {">vvvA", "vvv>A"},
    },
    '9': {
        '7': {"<<A"}, '8': {"<A"}, '9': {"A"},
        '4': {"<<vA", "v<<A"}, '5': {"<vA", "v<A"}, '6': {"vA"},
        '1': {"<<vvA", "vv<<A"}, '2': {"<vvA", "vv<A"}, '3': {"vvA"},
        '0': {"<vvvA", "vvv<A"}, 'A': {"vvvA"},
    },
}
var dirpad = map[rune]map[rune][]string{
    'A': {
        '^': {"<A"}, 'A': {"A"},
        '<': {"<v<A", "v<<A"}, 'v': {"<vA", "v<A"}, '>': {"vA"},
    },
    '^': {
        '^': {"A"}, 'A': {">A"},
        '<': {"v<A"}, 'v': {"vA"}, '>': {"v>A", ">vA"},
    },
    '<': {
        '^': {">^A"}, 'A': {">>^A", ">^>A"},
        '<': {"A"}, 'v': {">A"}, '>': {">>A"},
    },
    'v': {
        '^': {"^A"}, 'A': {">^A", "^>A"},
        '<': {"<A"}, 'v': {"A"}, '>': {">A"},
    },
    '>': {
        '^': {"^<A", "<^A"}, 'A': {"^A"},
        '<': {"<<A"}, 'v': {"<A"}, '>': {"A"},
    },
}


