package main

import (
    "aoc2024/data"
    "fmt"
    "strings"
    "strconv"
    "log"
)

// :P
var rules map[string]map[string]bool = make(map[string]map[string]bool)
var lists []string
var badlists []string

func main() {
    input := data.Get("data/day05.txt")

    startLists := false
    for i := range input {
        if input[i] == "" {
            startLists = true
            continue
        }

        if !startLists {
            kv := strings.Split(input[i], "|")

            _, key := rules[kv[0]]
            if key {
                rules[kv[0]][kv[1]] = true
            } else {
                rules[kv[0]] = make(map[string]bool)
                rules[kv[0]][kv[1]] = true
            }
        } else {
            lists = append(lists, input[i])
        }
    }

    fmt.Println("PART 1: ", part1())
    fmt.Println("PART 2: ", part2())
}

func part1() int {
    total := 0 

    for _, list := range lists {
        total += func() int {
            seq := strings.Split(list, ",")

            for in, val := range seq {
                for sub := in+1; sub < len(seq); sub++ {
                    _, key := rules[val][seq[sub]]
                    if !key {
                        badlists = append(badlists, list)
                        return 0
                    }
                }
            }
            num, err := strconv.Atoi(seq[len(seq) / 2])
            if err != nil {
                log.Fatal(err)
            }
            return num
        }()
    }
    return total
}

func part2() int {

    cmp := func(X, Y string) bool {
        return rules[X][Y]
    }

    swap := func(x, y int, seq *[]string) {
        holder := (*seq)[x]
        (*seq)[x] = (*seq)[y]
        (*seq)[y] = holder
    }

    total := 0 
    for _, list := range badlists {
        var slice []string = strings.Split(list, ",")
        ptr := &slice

        for i := 0; i < len(*ptr) - 1; i++ {
            swapped := false
            for j := 0; j < len(*ptr) - i - 1; j++ {
                if cmp((*ptr)[j], (*ptr)[j + 1]) {
                    swap(j, j + 1, ptr)
                    swapped = true
                }
            }

            if !swapped {
                break
            }
        }

        num, err := strconv.Atoi((*ptr)[len(*ptr) / 2])
        if err != nil {
            log.Fatal(err)
        }

        total += num 
    }
    return total
}
