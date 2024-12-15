package main

import (
    "bufio"
    "os"
    "log"
    "fmt"
    "strings"
    "strconv"
    "container/heap"
)


func main() {
    file, err := os.Open("data/day01.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(bufio.NewReader(file))

    var l1, l2 []int

    for scanner.Scan() {
        line := scanner.Text()

        split := strings.Split(line, "   ")

        i1, err := strconv.Atoi(split[0])
        if err != nil {
            log.Fatal(err)
        }
        i2, err := strconv.Atoi(split[1])
        if err != nil {
            log.Fatal(err)
        }

        l1 = append(l1, i1)
        l2 = append(l2, i2)
    } 

    fmt.Println("PART 1: ", part1(l1, l2))
    fmt.Println("PART 2: ", part2(l1, l2))

}

func part1(list1, list2 []int) int {
    heap1 := new(Heap)
    heap2 := new(Heap)

    heap.Init(heap1)
    heap.Init(heap2)

    for i := range list1{

        heap.Push(heap1, list1[i])
        heap.Push(heap2, list2[i])
    }

    sum := 0
    for range heap1.Len() {
        item1 := heap.Pop(heap1).(int)
        item2 := heap.Pop(heap2).(int)

        if item1 >= item2 { sum += item1 - item2 } 
        if item1 < item2  { sum += item2 - item1 }
    }
    return sum
}

func part2(list1, list2 []int) int {
    rhsMap := make(map[int]int)

    for i := range list2 {
        val, key := rhsMap[list2[i]]
        if key {
            rhsMap[list2[i]] = val + 1
        } else {
            rhsMap[list2[i]] = 1
        }
    }

    sum := 0
    for i := range list1 {
        val, key := rhsMap[list1[i]]
        if key {
            sum += list1[i] * val
        }
    }

    return sum
}

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
