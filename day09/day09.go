package main 


import ( 
    "aoc2024/data"
    "fmt"
    "strconv"
)

var input []string

func main() {
    input = data.Get("data/day09.txt")

    fmt.Println("PART 1:", part1())
    fmt.Println("PART 2:", part2())
}

func part1() int {
    disk := input[0]

    diskRep := make([]int, 0)
    elements := 0
    freeSpace := 0
    for i, str := range disk {
        num, _ := strconv.Atoi(string(str))
        if i % 2 == 0 {
            for range num {
                diskRep = append(diskRep, i / 2) 
                elements++
            }
        } else {
            for range num {
                diskRep = append(diskRep, -1)
                freeSpace++
            }
        }
    }

    for i := len(diskRep) - 1; i >= 0; i-- {
        num := diskRep[i]

        swapped := false
        for j := range elements {
            if diskRep[j] == -1 {
                diskRep[j] = num 
                diskRep[i] = -1
                swapped = true
                break
            }
        }
        if !swapped {
            break
        }
    }
    diskRep = diskRep[:elements]

    return func () int {
        total := 0
        for i, val := range diskRep {
            total += i * val
        }
        return total
    }()
}

func part2() int {
    disk := input[0]
    //1:1 rep
    diskRep := make([]int, 0)
    //the interval of each file
    files := make([][]int, 0)
    //the interval of each space
    spaces := make([][]int, 0)

    for i := range disk {
        if i%2 == 0 {
            id := i / 2
            fileSpace, _ := strconv.Atoi(string(disk[i]))
            files = append(files, []int{len(diskRep), len(diskRep)-1 + fileSpace})

            for j := 0; j < int(fileSpace); j++ {
                diskRep = append(diskRep, id)
            }
        } else {
            space, _ := strconv.Atoi(string(disk[i]))
            spaces = append(spaces, []int{len(diskRep), len(diskRep)-1 + space})

            for j := 0; j < int(space); j++ {
                diskRep = append(diskRep, -1)
            }
        }
    }

    for j := len(files) - 1; j >= 0; j-- {

        for i := 0; i < len(spaces); i++ {
            space := spaces[i]

            //ensure that the space is to the left of the file 
            // and that the space can fit in the file
            if space[1] < files[j][1] && space[1]-space[0] >= files[j][1]-files[j][0] {
                for k := space[0]; k <= space[0]+files[j][1]-files[j][0]; k++ {
                    diskRep[k] = j
                }

                for k := files[j][0]; k <= files[j][1]; k++ {
                    diskRep[k] = -1
                }

                space[0] = space[0] + files[j][1] - files[j][0] + 1
                break
            }
        }

    }

    total := 0
    for i := range diskRep {
        if diskRep[i] != -1 {
            total += i * diskRep[i]
        }
    }

    return total
}
