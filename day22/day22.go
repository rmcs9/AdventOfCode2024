package main 


import (
    "aoc2024/data"
    "fmt"
    "strconv"
    "log"
)

var input []string


func main() {
    input = data.Get("data/day22.txt")

    fmt.Println("PART 1:", part1()) 
    fmt.Println("PART 2:", part2())
}

func part1() int {
    total := 0
    for _, str := range input {
        secret, err := strconv.Atoi(str) 
        if err != nil {
            log.Fatal(err)
        }

        for range 2000 {
            secret = getNextSecret(secret)
        }

        total += secret
    }
    return total
}

func getNextSecret(secret int) int {
    mix   := func(s, m int) int { return s ^ m }
    prune := func(s int) int { return s % 16777216 }

    mixer := secret
    mixer  = mixer * 64 
    secret = mix(secret, mixer)
    secret = prune(secret)

    mixer  = secret / 32 
    secret = mix(secret, mixer)
    secret = prune(secret)

    mixer  = secret * 2048 
    secret = mix(secret, mixer)
    secret = prune(secret)
    return secret
}

func part2() int {
    monkeys := make([]monkey, len(input))
    for i, str := range input {
        secret, err := strconv.Atoi(str)
        if err != nil {
            log.Fatal(err)
        }
        changes := make([]change, 2000)
        changes[0] = change{secret % 10, 0}
        for j := 1; j < 2000; j++ {
            secret = getNextSecret(secret)
            changes[j] = change{secret % 10, (secret % 10) - changes[j - 1].price}
        }

        monkeys[i].secrets = changes
    }
    //map from sequence to bananas
    seqs := make(map[seq]int)
    // map to stop monkeys from contributing to a score more than once.
    // we can only sell on the first occurence of the sequence
    contributions := make(map[seq]map[int]bool)

    for i := 1; i < 1997; i++ {
        
        for m, monkey := range monkeys {
            //determine if the sequence is valid (its values vary between each index)
            valid := func() bool {
                prev := monkey.secrets[i].delta
                for j := i + 1; j < i + 4; j++ {
                    current := monkey.secrets[j].delta 
                    if prev == current {
                        return false
                    }
                    prev = current
                }
                return true
            }()

            // if a seq is valid and its score is greater than 0, 
            if valid && monkey.secrets[i + 3].price > 0 {
                //create its sequence key
                ks := make([]int, 4)
                for j := range 4 {
                    ks[j] = monkey.secrets[i + j].delta
                }
                key := seq{ks[0], ks[1], ks[2], ks[3]}

                if _, k := contributions[key]; !k {
                    contributions[key] = make(map[int]bool)
                }

                //add its score to the map
                if val := contributions[key][m]; !val {
                    seqs[key] += monkey.secrets[i + 3].price
                    contributions[key][m] = true
                }

            }
        }
    }

    //take the highest score out of all the sequences
    bananas := 0
    for _, val := range seqs {
        if val > bananas {
            bananas = val
        }
    }
    return bananas
}

type change struct {
    price, delta int
}

type monkey struct {
    secrets []change
}

// cant use slices as keys in maps... cringe
type seq struct {
    p1, p2, p3, p4 int
}
