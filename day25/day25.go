package main 


import (
    "aoc2024/data"
    "fmt"
)

var keys  = make([][]int, 0)
var locks = make([][]int, 0)
var lockSpace = 5

func main() {
    input := data.Get("data/day25.txt")
    input = append(input, "")

    base := "" 
    for range input[0] { base += "#" }

    lockkey := make([]string, 0)
    for _, line := range input {
        if line == "" {
            cols := make([]int, len(lockkey[0]))
            for i := 1; i < len(lockkey) - 1; i++ {
                for j, c := range lockkey[i] {
                    if c == '#' {
                        cols[j] += 1
                    }
                }
            }
            if lockkey[0] == base {
                locks = append(locks, cols)
            } else {
                keys = append(keys, cols)
            }
            lockkey = make([]string, 0)
            continue
        }

        lockkey = append(lockkey, line)
    }

    fmt.Println("PART 1:", part1()) 
}

func part1() int {
    
    pairs := make(map[string]bool)
    for k, key := range keys {
        for l, lock := range locks {
            
            fits := true
            for i := 0; i < len(key); i++ {
                if key[i] + lock[i] > lockSpace {
                    fits = false 
                    break
                }
            }

            if fits {
                pairs[fmt.Sprintf("%d %d", k, l)] = true
            }
        }
    }

    return len(pairs)
}
