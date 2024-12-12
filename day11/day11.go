package main 


import (
	"aoc2024/data"
	"fmt"
	"container/list"
	"strings"
	"strconv"
)

var nums []string

func main() {
	input := data.Get("data/day11.txt")
	
	numsStr := strings.Split(input[0], " ")
	for _, val := range numsStr {
		nums = append(nums, val)
	}
	
	fmt.Println("PART 1:", part1())
	fmt.Println("PART 2:", part2())
}

func part1() int {
	stones := list.New()

	for _, val := range nums {
		stones.PushBack(val)
	}

	for range 25 {
		for s := stones.Front(); s != nil; s = s.Next() {
			if s.Value.(string) == "0" {
				s.Value = "1"
			} else if len(s.Value.(string)) % 2 == 0 {
				str := s.Value.(string)
				half := str[:len(str)/2]
				stones.InsertBefore(half, s)
				half2 := str[len(str)/2:]
				num, _ := strconv.Atoi(half2)

				if num == 0 {
					s.Value = "0"
				} else {
					for half2[0] == '0' {
						half2 = half2[1:]
					}
					s.Value = half2
				}
			} else {
				num, _ := strconv.Atoi(s.Value.(string))
				num = num * 2024

				s.Value = strconv.Itoa(num)
			}
		}
	}
	return stones.Len() 
}

func part2() int {
	stones := make(map[string]int) 

	for _, val := range nums {
		stones[val] = 1
	}

	for range 75 {
		newStones := make(map[string]int)
		for key, val := range stones {
			if key == "0" {
				c, _:= newStones["1"] 
				newStones["1"] = c + val

			} else if len(key) % 2 == 0 {
				str := key
				half := str[:len(str)/2]
				half2 := str[len(str)/2:]
				
				if num, _ := strconv.Atoi(half2); num == 0 {
					half2 = "0"
				} else {
					for half2[0] == '0' {
						half2 = half2[1:]
					}
				}

				c, _ := newStones[half] 
				newStones[half] = c + val
				c, _ = newStones[half2] 
				newStones[half2] = c + val

			} else {
				num, _ := strconv.Atoi(key)
				num = num * 2024

				s := strconv.Itoa(num)
				c, _ := newStones[s] 
				newStones[s] = c + val
			}
		}
		stones = newStones
	}
	total := 0
	for _, val := range stones {
		total += val
	}
	return total
}
