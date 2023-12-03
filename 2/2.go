package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Subset struct {
    red int
    green int
    blue int
}


func max(x int, y int) int {
    if x > y {
        return x
    }
    return y
}

const RED_REQUIRED = 12
const GREEN_REQUIRED = 13
const BLUE_REQUIRED = 14 

func (s Subset) isPossibleGame() bool {
    return s.red <= RED_REQUIRED && s.green <= GREEN_REQUIRED && s.blue <= BLUE_REQUIRED
}

func (s Subset) getPower() int {
    return s.red * s.green * s.blue 
}

func main() {
	start := time.Now()
	file, err := os.Open("input")

	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()

    m := make(map[string]Subset)
    

    sumIds := 0
	scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        splitByColon := strings.Split(line, ":")
        gameStr, subsetsList := splitByColon[0], splitByColon[1]

        gameSplit := strings.Split(gameStr, " ")
        gameId := gameSplit[1]

        subsets := strings.Split(subsetsList, ";")

        for _, subset := range subsets {
            for _, cube := range strings.Split(subset, ",") {
                if _, ok := m[gameId]; !ok {
                    m[gameId] = Subset{0, 0, 0}
                }

                splitCube := strings.Split(strings.TrimSpace(cube), " ")
                count, countErr  := strconv.Atoi(splitCube[0])
                if countErr != nil {
                    fmt.Println("error converting count to int")
                    return
                }
                color := splitCube[1]

                switch color {
                    case "red":
                        m[gameId] = Subset{max(m[gameId].red, count), m[gameId].green, m[gameId].blue }
                    case "green":
                        m[gameId] = Subset{m[gameId].red, max(m[gameId].green, count), m[gameId].blue }
                    case "blue":
                        m[gameId] = Subset{m[gameId].red, m[gameId].green, max(m[gameId].blue, count)}
                }
            }
        }
    }

    sumPower := 0
    for k, v := range m {
        if v.isPossibleGame() {
            id, errId := strconv.Atoi(k)
            if errId != nil {
                fmt.Println("error converting id to int")
                return
            }
            sumIds += id
        }
        sumPower += v.getPower()
    }

    fmt.Println("sum: ", sumIds)
    fmt.Println("power: ", sumPower)

    if err := scanner.Err(); err != nil {
		fmt.Println("error scanning file")
	}


	fmt.Println("Running time", time.Since(start))
}
