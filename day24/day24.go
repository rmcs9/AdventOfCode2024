package main 


import (
    "aoc2024/data"
    "fmt"
    "sync"
    "strings"
    "strconv"
    "sort"
    "time"
)

// from sig -> val
var signals = sync.Map{}
var gates = make([]gate, 0)
var gatesMap = make(map[string][]string)

type gate struct {
    x, y, r string 
    log func(int, int) int 
}

func main() {
    input := data.Get("data/day24.txt")

    or := func(x,y int) int {
        return x | y
    }

    and := func(x,y int) int {
        return x & y
    }

    xor := func(x,y int) int {
        return x ^ y
    }

    gs := false 
    xstr, ystr := "", ""
    for _, line := range input {
        if line == "" {
            gs = true 
            continue
        }

        if !gs {
            col := strings.Index(line, ":")
            sig, val := line[:col], line[col + 2:]
            nval, _ := strconv.Atoi(val)

            signals.Store(sig, nval)
            if line[0] == 'x' {
                xstr += val 
            } else {
                ystr += val
            }
        } else {
            gsplit := strings.Split(line, " ")
            
            var g gate
            g.x, g.y, g.r = gsplit[0], gsplit[2], gsplit[4]
            if gsplit[1] == "XOR" { g.log = xor } 
            if gsplit[1] == "OR"  { g.log = or  }
            if gsplit[1] == "AND" { g.log = and }
            gates = append(gates, g)
            gatesMap[g.r] = []string{ g.x, g.y, gsplit[1] }
        }
    }

    fmt.Println("PART 1:", part1()) 
    fmt.Println("PART 2:", part2())
}


func part1() int {
    var wg sync.WaitGroup
    for _, g := range gates {
        wg.Add(1)
        go func(gate gate) {
            defer wg.Done()
            xval, ok := signals.Load(gate.x) 
            for !ok {
                time.Sleep(1 * time.Microsecond)
                xval, ok = signals.Load(gate.x)
            }

            yval, ok := signals.Load(gate.y) 
            for !ok {
                time.Sleep(1 * time.Microsecond)
                yval, ok = signals.Load(gate.y) 
            }

            res := gate.log(xval.(int), yval.(int))
            signals.Store(gate.r, res)
        }(g)
    }

    wg.Wait()
    zs := make([]string, 0) 
    signals.Range(func(key, value interface{}) bool {
        keystr := key.(string) 
        if keystr[0] == 'z' {
            zs = append(zs, keystr)
        }
        return true 
    })
    sort.Strings(zs)

    bstr := "" 
    for i := len(zs) - 1; i >= 0; i-- {
        z := zs[i]
        val, _ := signals.Load(z) 
        valstr := strconv.Itoa(val.(int))
        bstr += valstr
    }

    ret, _ := strconv.ParseInt(bstr, 2, 64)
    return int(ret)
}

func part2() string {
    ඞ := make(map[string]bool)
    
    for ඞ_ඞ, ඞ_ඞ_ඞ := range gatesMap {
        if ඞ_ඞ[0] == 'z' && ඞ_ඞ != "z00" && ඞ_ඞ != "z45" {
            if ඞ_ඞ_ඞ[2] != "XOR" { ඞ[ඞ_ඞ] = true }
        }

        if ඞ_ඞ_ඞ[2] == "XOR" && ඞ_ඞ_ඞ[0][0] != 'x' && ඞ_ඞ_ඞ[1][0] != 'x' && ඞ_ඞ != "z01" {
            if ඞ_ඞ[0] != 'z' { ඞ[ඞ_ඞ] = true }
            if gatesMap[ඞ_ඞ_ඞ[0]][2] == "AND" { ඞ[ඞ_ඞ_ඞ[0]] = true }
            if gatesMap[ඞ_ඞ_ඞ[1]][2] == "AND" { ඞ[ඞ_ඞ_ඞ[1]] = true }
        }

        if ඞ_ඞ_ඞ[2] == "OR" {
            if gatesMap[ඞ_ඞ_ඞ[0]][2] == "XOR" { ඞ[ඞ_ඞ_ඞ[0]] = true }
            if gatesMap[ඞ_ඞ_ඞ[1]][2] == "XOR" { ඞ[ඞ_ඞ_ඞ[1]] = true }
        }
    }

    ඞඞ := make([]string, 0)
    for key := range ඞ {
        ඞඞ = append(ඞඞ, key)
    }
    sort.Strings(ඞඞ)
    ඞඞඞ := ""
    for _, ඞඞඞඞ := range ඞඞ {
        ඞඞඞ += ඞඞඞඞ + ","
    }
    return ඞඞඞ[:len(ඞඞඞ) - 1] 
}
