package main 


import (
    "aoc2024/data"
    "fmt"
    "strings"
    "sort"
)

var net map[string]map[string]bool

func main() {
    input := data.Get("data/day23.txt")

    net = make(map[string]map[string]bool)

    for _, conn := range input {
        clients := strings.Split(conn, "-")

        if _, key := net[clients[0]]; !key {
            net[clients[0]] = make(map[string]bool)
        }

        if _, key := net[clients[1]]; !key {
            net[clients[1]] = make(map[string]bool)
        }

        net[clients[0]][clients[1]] = true
        net[clients[1]][clients[0]] = true
    }

    fmt.Println("PART 1:", part1()) 
    fmt.Println("PART 2:", part2())
}


func part1() int {
    // for every client that starts with a t, intersect all of its connections by groups of 3 and determine if all 3 are shared
    explored := make(map[string]bool)
    total := 0
    for client, conns := range net {
        if client[0] == 't' {
            for set1key := range conns {
                for set2key := range conns {
                    if set1key == set2key {
                        continue
                    }

                    total += func() int {
                        key1 := net[set2key][set1key] 
                        key2 := net[set1key][set2key] 

                        exKeys := []string{client, set1key, set2key}
                        sort.Strings(exKeys)
                        exkey := exKeys[0] + exKeys[1] + exKeys[2]
                        if ex := explored[exkey]; !ex && key1 && key2 {
                            explored[exkey] = true
                            return 1
                        }
                        return 0
                    }()
                }
            }
        }
    }
    return total
}



// for every client the goal is to determine the largest network within that client 
// then take the largest network out of all the clients
//
// every client has a set max amount of connections, this can be observed in the graph
// therfore the maximum size of the largest network can be at most the max amount of connections every client has (13 in my input)
//
// this makes it easy to break up the network finding by client
func part2() string {
    // current leader among the clients
    leader := make([]string, 0)
    for client, conns := range net {
        // current leader within the client
        clientLeader := make([]string, 0)
        for s1key := range conns {
            // the group that is able to be formed with { client, s1key }...
            // any node that is connected to the client, s1, and any subsequently added keys, is brought into the group
            connGroup := []string{client, s1key}
            for s2key := range conns {
                if s1key == s2key { continue }
                
                //check if s2 is connected to s1
                if inS1 := net[s1key][s2key]; inS1 {
                    //then check if s2 is connected to everything else in the group
                    inGroup := true 
                    for _, key := range connGroup {
                        if inG := net[key][s2key]; !inG {
                            inGroup = false 
                            break
                        }
                    }

                    // if it is, add it to the group
                    if inGroup {
                        connGroup = append(connGroup, s2key)
                    }
                }
            }
            // if this conn group beats out the leader, 
            // it takes the leaders place
            if len(connGroup) > len(clientLeader) {
                clientLeader = connGroup
            }
        }

        // if this clients leader beats out the the current leader, 
        // it takes the current leaders place
        if len(clientLeader) > len(leader) {
            leader = clientLeader
        }
    }

    // sort the winner and format 
    sort.Strings(leader) 
    ret := ""
    for _, val := range leader {
        ret += val + ","
    }
    ret = ret[:len(ret) - 1]
    return ret
}
