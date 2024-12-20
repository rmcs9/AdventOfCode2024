package main 


import (
    "aoc2024/data"
    "fmt"
    "strings"
)

var patterns map[string]bool 
var maxP int = 0
var designs []string
var memo map[string]int

func main() {
    input := data.Get("data/day19.txt")

    patterns = func() map[string]bool {
        pSplit := strings.Split(input[0], ",")

        ret := make(map[string]bool)
        for _, val := range pSplit {
            val = strings.TrimSpace(val)
            if len(val) > maxP { maxP = len(val) }
            ret[val] = true
        }
        return ret
    }()

    designs = input[2:]

    fmt.Println("PART 1:", part1()) 
    fmt.Println("PART 2:", part2())
}


func part1() int {
    total := 0
    for _, design := range designs {
        if matcher(design) {
            total += 1 
        }
    }
    return total
}

func matcher(design string) bool {
    //work your way down from the biggest pattern seeing if any of the ranges match
    for i := maxP; i > 0; i-- {
        // if there is still more design to eval, recur across the pattern
        if len(design) > i {
            if patterns[design[:i]] {
                if matcher(design[i:]) { return true }
            }
        //base case, if you have found the smallest possible chunk of the design,
        // evaluate it and return 
        } else if len(design) == i {
            if patterns[design] { return true }
        }
    }
    return false
}

func part2() int {
    memo = make(map[string]int)
    total := 0 
    for _, design := range designs {
        total += matcherp2(design)
    }
    return total
}

// pretty much the same as part 1 but now you cant return early if a single case matches
// all possible cases must be recorded and summed
// memoization makes this a lot faster
func matcherp2(design string) int {
    total := 0
    for i := maxP; i > 0; i-- {
        if len(design) > i {
            if patterns[design[:i]] {
                if count, key := memo[design[i:]]; key {
                    total += count
                } else {
                    subtotal := total
                    total += matcherp2(design[i:])
                    memo[design[i:]] = total - subtotal
                }
            }
        } else if len(design) == i {
            if patterns[design] { total += 1 }
        }
    }
    return total
}
