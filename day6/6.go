package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Race struct {
    time float64 
    distance float64 
}


func computeHoldTime(race *Race) (float64, float64) {
    T := race.time
    d := race.distance + 1
    t1 := (T - math.Sqrt(T*T - 4*d)) / 2
    t2 := (T + math.Sqrt(T*T - 4*d)) / 2
    t1 = math.Ceil(math.Abs(t1))
    t2 = math.Floor(math.Abs(t2))
    return t1, t2
}

func main() {
	start := time.Now()

	file, err := os.Open("input")

	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

    var races []*Race
    lineId := 0

    p2Race := &Race{time: 0, distance: 0}

    waysToWin := 1

    for scanner.Scan() {
        line := scanner.Text()

        splitByColon := strings.Split(line, ":")
        re := regexp.MustCompile(`\s+`)
        removeSpacesReg := regexp.MustCompile(`\s`)

        splitBySpace := strings.Split(strings.TrimSpace(re.ReplaceAllString(splitByColon[1], " ")), " ") 


        if(lineId == 0) {
            time := re.ReplaceAllString(splitByColon[1], "")
            num, err := strconv.Atoi(time)
            if err != nil {
                fmt.Println("error converting string to int")
            }
            p2Race.time = float64(num)
        } else {
            distance := removeSpacesReg.ReplaceAllString(splitByColon[1], "")
            num, err := strconv.Atoi(distance)
            if err != nil {
                fmt.Println("error converting string to int")
            }
            p2Race.distance = float64(num)
        }

        for i, val := range splitBySpace {
            num, err := strconv.Atoi(val)
            if err != nil {
                fmt.Println("error converting string to int")
            }
           if(lineId == 0) {
               races = append(races, &Race{time: float64(num), distance: 0}) 
           } else {
               races[i].distance = float64(num)
           }
        }
        lineId++
    }

    for _, race := range races {
        // v = d / t
        // here is d determined by time hold = t
        // T is total time, t is time to hold
        // -b -+ sqrt(b^2 - 4ac) / 2a = get roots of quadratic formula
        t1, t2 := computeHoldTime(race)
        waysToWin *= int(t2) - int(t1) + 1
    }
    fmt.Println(waysToWin)

    // part two
    t1, t2 := computeHoldTime(p2Race)
    waysToWinP2 := int(t2) - int(t1) + 1
    fmt.Println("partTwo", waysToWinP2)

    fmt.Println("running time", time.Since(start))
}
