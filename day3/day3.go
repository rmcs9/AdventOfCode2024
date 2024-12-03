package main

import ( 
	"aoc2024/data"
	"fmt"
	"strings"
	"strconv"
	"regexp"
)

func main() {
	input := data.Get("data/day3.txt")

	fmt.Println("PART 1: ", part1(input))
	fmt.Println("PART 2: ", part2(input))
}

func part1(input []string) int {

	mults := 0
	for i := range input {
		funcs := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		exe := funcs.FindAllString(input[i], -1)

		for k := range exe {
			mults += func(mul string) int {
				mul, _ = strings.CutPrefix(mul, "mul(")
				mul, _ = strings.CutSuffix(mul, ")")

				nums := strings.Split(mul, ",")

				f1, _ := strconv.Atoi(nums[0])
				f2, _ := strconv.Atoi(nums[1])

				return f1 * f2
			}(exe[k])
		}

	}
	return mults
}

func part2(input []string) int {
	
	mults := 0
	enabled := true
	for i := range input {
		funcs := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|(do\(\))|(don't\(\))`)
		exe := funcs.FindAllString(input[i], -1)
		
		for k := range exe {
			if exe[k] == "don't()" {
				enabled = false
			} else if exe[k] == "do()" {
				enabled = true 
			} else if enabled {
				mults += func(mul string) int {
					mul, _ = strings.CutPrefix(mul, "mul(")
					mul, _ = strings.CutSuffix(mul, ")")

					nums := strings.Split(mul, ",")

					f1, _ := strconv.Atoi(nums[0])
					f2, _ := strconv.Atoi(nums[1])

					return f1 * f2
				}(exe[k])
			}
		}
	}
	return mults
}
