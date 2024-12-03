package main


import ( 
	"aoc2024/data"
	"fmt"
	"strconv"
	"log"
	"strings"
)

func main() {
	input := data.Get("data/day2.txt")

	fmt.Println("PART 1: ", part1(input))
	fmt.Println("PART 2: ", part2(input))
}

var safe = func(list []int) bool {

	prev := list[0]
	var delta bool

	for k := 1; k < len(list); k++ {

		if abs(prev - list[k]) > 3 || abs(prev - list[k]) < 1 {
			return false
		}

		newdelta := prev < list[k]

		if k == 1 {
			delta = newdelta
		} else if delta != newdelta {
			return false
		}
		prev = list[k]
	}
	return true

}

var abs = func (i int) int {
	if i < 0 { i = -i } 
	return i 
}

func part1(input []string) int {

	saferanges := 0

	for i := range input {
		nums := func(in []string) []int {
			var n []int
			for i := range in {
				num, err := strconv.Atoi(in[i])
				if err != nil {
					log.Fatal(err)
				}
				n = append(n, num)
			}
			return n
		}(strings.Split(input[i], " "))

		if safe(nums) { saferanges++ }
	}
	return saferanges
}

func part2(input []string) int {
	saferanges := 0

	for i := range input {
		nums := func(in []string) []int {
			var n []int
			for i := range in {
				num, err := strconv.Atoi(in[i])
				if err != nil {
					log.Fatal(err)
				}
				n = append(n, num)
			}
			return n
		}(strings.Split(input[i], " "))

		if safe(nums) {
			saferanges++ 
			continue
		}

		if safe(nums[1:]) {
			saferanges++
			continue 
		}

		sec := make([]int, 0)
		sec = append(sec, nums[0])
		sec = append(sec, nums[2:]...)

		if safe(sec) {
			saferanges++ 
			continue
		}

		//:P
		if func(list []int) bool {
			prev := list[1]
			delta := list[1] - list[0] > 0

			for k := 2; k < len(list); k++ {
				if (abs(prev - list[k]) < 1 || abs(prev - list[k]) > 3) ||
				(delta && list[k] - prev < 0) ||
				(!delta && list[k] - prev > 0) {
					var prevslice []int 
					var kslice []int 

					prevslice = append(prevslice, list[:k-1]...)
					prevslice = append(prevslice, list[k:]...)

					kslice = append(kslice, list[:k]...)
					kslice = append(kslice, list[k+1:]...)


					if !safe(prevslice) && !safe(kslice) {
						return false
					}
				}
				prev = list[k]
			}
			return true
		} (nums) {
			saferanges++
		}
	}	
	return saferanges
}
