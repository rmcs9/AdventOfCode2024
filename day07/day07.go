package main 


import ( 
	"aoc2024/data"
	"fmt"
	"strings"
	"strconv"
	"log"
	"github.com/mowshon/iterium"
)

var input []string

func main() {
	input = data.Get("data/day07.txt")

	fmt.Println("PART 1:", part1())
	fmt.Println("PART 2:", part2())
}

var getArgs = func (args string) []int {
	strings := strings.Split(strings.TrimSpace(args), " ")

	nums := make([]int, len(strings))
	for i, str := range strings {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		nums[i] = num
	}
	return nums
}

func part1() int {
	total := 0
	for _, eq := range input {
		eqSplit := strings.Split(eq, ":")

		res, args := eqSplit[0], eqSplit[1]

		resNum, _ := strconv.Atoi(res)

		argList := getArgs(args)
		
		permSlice, _ := iterium.Product([]string { "+", "*" }, len(argList) - 1).Slice()
		
		total += func() int {
			for _, perm := range permSlice {
				sum := argList[0] 

				for i, sign := range perm {
					if sign == "+" {
						sum += argList[i + 1]
					} else if sign == "*"{
						sum *= argList[i + 1]
					}

					if sum > resNum { break }
				}

				if sum == resNum {
					return resNum
				}
			}
			return 0
		}()

	}
	return total
}

func part2() int {	
	total := 0
	ch := make(chan int)
	for  _, eq := range input {
		go p2worker(ch, eq)
	}

	for range input {
		total += <-ch
	}
	close(ch)
	return total
}

func p2worker(ch chan int, eq string) {
	eqSplit := strings.Split(eq, ":")
	res, args := eqSplit[0], eqSplit[1]
	resNum, _ := strconv.Atoi(res)
	argList := getArgs(args)
	permSlice, _ := iterium.Product([]string { "+", "*", "||" }, len(argList) - 1).Slice()
	
	ch <- func() int {
		for _, perm := range permSlice {
			sum := argList[0] 

			for i, sign := range perm {
				if sign == "+" {
					sum += argList[i + 1]
				} else if sign == "*"{
					sum *= argList[i + 1]
				} else if sign == "||" {
					sum, _ = strconv.Atoi(strconv.Itoa(sum) + strconv.Itoa(argList[i + 1]))
				}

				if sum > resNum { break }
			}

			if sum == resNum {
				return resNum
			}
		}
		return 0
	}()
}
