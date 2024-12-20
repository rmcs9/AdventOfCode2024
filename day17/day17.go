package main 


import (
    "aoc2024/data"
    "fmt"
    "strings"
    "strconv"
    "log"
    "math"
)

var instructions []int
var regA, regB, regC int
var instStr string

func main() {
    input := data.Get("data/day17.txt")
    instStr = input[4][strings.Index(input[4], ":") + 2:]
    regA, _ = strconv.Atoi(input[0][strings.Index(input[0], ":") + 2:])
    regB, _ = strconv.Atoi(input[1][strings.Index(input[1], ":") + 2:])
    regC, _ = strconv.Atoi(input[2][strings.Index(input[2], ":") + 2:])

    instructions = func() []int {
        list := input[4][strings.Index(input[4], ":") + 2:]
        splitList := strings.Split(list, ",")

        inst := make([]int, 0)
        for _, val := range splitList {
            num, err := strconv.Atoi(val)
            if err != nil {
                log.Fatal(err)
            }
            inst = append(inst, num)
        }
        return inst
    }()

    fmt.Println("PART 1:", part1(-1)) 
    fmt.Println("PART 2:", part2())
}

func part1(a int) string {
    PC := 0 
    out := ""

    registers := make(map[int]int)
    if a != -1 {
        registers[4] = a 
    } else {
        registers[4] = regA
    }
    registers[5] = regB 
    registers[6] = regC

    for PC < len(instructions) {
        //adv 
        if instructions[PC] == 0 {
            denom, key := registers[instructions[PC + 1]]
            if !key {
                denom = instructions[PC + 1] 
            }
            denom = int(math.Pow(2, float64(denom)))
            registers[4] = int(registers[4] / denom) 
        // bxl 
        } else if instructions[PC] == 1 {
            registers[5] = registers[5] ^ instructions[PC + 1]
        // bst
        } else if instructions[PC] == 2 {
            combo, key := registers[instructions[PC + 1]]
            if !key {
                combo = instructions[PC + 1]
            }
            registers[5] = combo % 8
        //jnz
        } else if instructions[PC] == 3 {
            if registers[4] != 0 {
                PC = instructions[PC + 1]
                continue
            }
        // bxc
        } else if instructions[PC] == 4 {
            registers[5] = registers[5] ^ registers[6]
        // out
        } else if instructions[PC] == 5 {
            combo, key := registers[instructions[PC + 1]]
            if !key {
                combo = instructions[PC + 1]
            }
            combo = combo % 8
            if out == "" {
                out += fmt.Sprint(combo)
            } else {
                out += fmt.Sprint(",", combo)
            }
        // bdv
        } else if instructions[PC] == 6 {
            denom, key := registers[instructions[PC + 1]] 
            if !key {
                denom = instructions[PC + 1]
            }
            denom = int(math.Pow(2, float64(denom))) 
            registers[5] = int(registers[4] / denom)
        // cdv
        } else if instructions[PC] == 7 {
            denom, key := registers[instructions[PC + 1]] 
            if !key {
                denom = instructions[PC + 1]
            }
            denom = int(math.Pow(2, float64(denom)))
            registers[6] = int(registers[4] / denom)
        }
        PC += 2
    }
    return out
}

func part2() int {
    return p2(instStr)
}

func p2(target string) int {
    val := 0
    if len(target) != 1 {
        val = 8 * p2(target[2:])
    }

    for target != part1(val) {
       val++
    }
    return val
}
