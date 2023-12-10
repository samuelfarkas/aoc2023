package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
) 

type MapItem struct {
    src int
    dst int
    length int
}

type Range struct {
    to string
    values []*MapItem
}

func main() {
	start := time.Now()

	file, err := os.Open("input")

	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

    var initialSeeds []int
    var locations []int
    ranges := make(map[string]*Range)
    

    lineId := -2 
    mapIdx := []string{"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}

    for _, idx := range mapIdx {
        ranges[idx] = &Range{to: "", values: make([]*MapItem, 0)}
    }

    var seen []bool
    for scanner.Scan() {
        line := scanner.Text()
        // skip empty lines
        if line == "" {
            continue
        }
        if(lineId == -2) {
            splitByColon := strings.Split(line, ":")
            splitBySpace := strings.Split(splitByColon[1], " ")
            for _, seed := range splitBySpace {
                seed = strings.TrimSpace(seed)
                if(seed != "") {
                    seedInt, err := strconv.Atoi(seed)
                    if err != nil {
                        fmt.Println("error converting string to int")
                    }
                    initialSeeds = append(initialSeeds, seedInt)
                }
            }

            seen = make([]bool, len(locations))
            lineId++
            continue
        } 


        if(unicode.IsDigit(rune(line[0]))) {
            splitBySpace := strings.Split(line, " ")
            var numbers [3]int
            for i, val := range splitBySpace {
                number, err := strconv.Atoi(val)
                if err != nil {
                    fmt.Println("error converting string to int")
                }
                numbers[i] = number
            }

            // mapName := mapIdx[lineId]
            if(lineId < len(mapIdx)) {
                destination := numbers[0]
                source := numbers[1]
                length := numbers[2]
                // ranges[mapName].to = mapIdx[lineId + 1]
                // ranges[mapName].values = append(ranges[mapName].values, &MapItem{src: source, dst: destination, length: length})

                tempLocations := make([]int, len(locations))
                for i, loc := range locations {
                        tempLocations[i] = loc
                        if (loc >= source && loc < source + length) {
                             if(!seen[i]) {
                                seen[i] = true
                                result := (destination - source) + loc 
                                tempLocations[i] = result
                            }                        

                    }                     
                }
                locations = tempLocations
            }
        } else {
            seen = make([]bool, len(locations))
            lineId++
        }
    }

    fmt.Println("final locs", locations)

    low := locations[0]
    for _, loc := range locations {
        if(loc < low) {
            low = loc
        }
    }
    fmt.Println("lowest location", low)
    // for _, key := range mapIdx {
    //     rn := ranges[key]
    //     fmt.Println(key, rn.to)
    //     for _, item := range rn.values {
    //         fmt.Println(item.dst, item.src, item.length)
    //     }
    // }
    fmt.Println("time elapsed: ", time.Since(start))
}
