package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

var MOVES map[byte]int = map[byte]int{'L': 0, 'R': 1}

func walk(m map[string][]string, instructions []byte, startingNode string) int {
    nextMove := startingNode
    steps := 0
     
    for nextMove[2] != 'Z' {
        nextMove = m[nextMove][MOVES[instructions[steps % len(instructions)]]]
        steps += 1
    }
    return steps 
}

// greatest common divisor
func GCD(a, b int) int {
    for b != 0 {
        t := b
        b = a % b
        a = t
    }
    return a
}
  
// least common multiple
func LCM(a, b int) int {
    return (a * b) / GCD(a, b)
}
  
func LCMofArray(numbers []int) int {
    lcm := numbers[0]
    n := len(numbers)
   
    for i := 1; i < n; i++ {
        lcm = LCM(lcm, numbers[i])
    }
   
    return lcm
} 

func main() {
    start := time.Now()
	file, err := os.Open("input")

	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

    var instructions []byte

    m := make(map[string][]string)
    instructionsDone := false
    var startingNodes []string

    for scanner.Scan() {
        line := scanner.Text()

        if line == "" {
            instructionsDone = true
            continue
        }

        if !instructionsDone {
            instructions = append(instructions, line...)
        } else {
            // delete spaces
            re := regexp.MustCompile(`\s+`)
            line = re.ReplaceAllString(line, "")
            splitByEqual := strings.Split(line, "=")
            key, values := splitByEqual[0], splitByEqual[1]
            
            valuesSplitByComma := strings.Split(values, ",")
            left := valuesSplitByComma[0][1:]
            right := valuesSplitByComma[1][:3]
            
            m[key] = []string{left, right}

            if key[2] == 'A' {
                startingNodes = append(startingNodes, key)
            }
        }
    }

    // part 1
    // part1 := func() {
    //     steps := 0
    //     nextMove := "AAA"
    //     for nextMove != "ZZZ" {
    //         nextMove = m[nextMove][MOVES[instructions[steps % len(instructions)]]]
    //         steps += 1
    //     }
    //     fmt.Println("Part 1: ", steps)
    //     fmt.Println("Time elapsed: ", time.Since(start))
    // }
    // part1()


    // part 2
    var stepsTaken []int
    for _, startingNode := range startingNodes {
        steps := walk(m, instructions, startingNode)
        stepsTaken = append(stepsTaken, steps)
    }
    
    // reddit yonk, number theory - in this problem we want to find the least steps needed to satisfy all cycles that are walked simultaneously
    // if some paths are shorter than LCM, then they will be repeated before all paths are done
    fmt.Println("Part 2: ", LCMofArray(stepsTaken))


    fmt.Println("Time elapsed: ", time.Since(start))
}
