package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)


func allZero(arr []int) bool {
    for v := range arr {
        if v != 0 {
            return false 
        }
    }
    return true 
}

func predictForward(history []int) int {
    if allZero(history) {
        return 0 
    }

    sequence := make([]int, 0) 
    for i:= 0; i < len(history) - 1; i++ {
        sequence = append(sequence, history[i+1] - history[i])
    }

    return predictForward(sequence) + history[len(history) - 1] 
}

func predictBackwards(history []int) int {
    if allZero(history) {
        return 0 
    }

    sequence := make([]int, len(history) - 1) 

    for i := len(history) - 1; i > 0; i-- {
        sequence[i - 1] = history[i - 1] - history[i]
    }

    return predictBackwards(sequence) + history[0]
}
    



func main() {
    start := time.Now()
    file, err := os.Open("input")

    if err != nil {
        fmt.Println("error opening file")
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    sum := 0
    sum_pt2 := 0
    for scanner.Scan() {
        line := scanner.Text()
        splitBySpace := strings.Split(line, " ")
        history := make([]int, 0)
        for v := range splitBySpace {
            num, err := strconv.Atoi(splitBySpace[v])
            if err != nil {
                fmt.Println("error converting bid to int")
            }
            history = append(history, num)
        }
        sum += predictForward(history)
        sum_pt2 += predictBackwards(history)
    }

    fmt.Println("part 1:", sum)
    fmt.Println("part 2:", sum_pt2)
    

    fmt.Println("running time:", time.Since(start))
}
