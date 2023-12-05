package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {
	start := time.Now()

	file, err := os.Open("input")

	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
    sum:= 0

    gameMatches := make(map[int]int)

	for scanner.Scan() {
        winningNumbers := make(map[string]bool)
		line := scanner.Text()
        splitByColon := strings.Split(line, ":")
        gameId, scratchNumbers := splitByColon[0], splitByColon[1] 
        re := regexp.MustCompile(`\s+`)
        gameId = re.ReplaceAllString(gameId, " ") 
        gameId = strings.Split(gameId, " ")[1]

        gameIdAsInt, err := strconv.Atoi(gameId)

        if err != nil {
            fmt.Println("error converting string to int")
        }
        gameMatches[gameIdAsInt] = 0

        splitByPipe := strings.Split(scratchNumbers, "|")
        winning, guessed := splitByPipe[0], splitByPipe[1] 

        number := "" 
        for _, char := range winning {
            if unicode.IsDigit(char) {
                number += string(char)
            }
            if(number != "" && !unicode.IsDigit(char)) {
                winningNumbers[number] = true
                number = ""
            }
        }
        
        rowPoints := 0 
        num := "" 

        for i, char := range guessed {
            if unicode.IsDigit(char) {
                num += string(char)
            }
            if (num != "" && !unicode.IsDigit(char)) || (num != "" && unicode.IsDigit(char) && i == len(guessed) - 1) {
                if(winningNumbers[num]) {
                    gameMatches[gameIdAsInt] += 1
                    if(rowPoints == 0) {
                        rowPoints += 1
                    } else {
                        rowPoints *= 2
                    }
                }
                num = ""
            }
        }

        sum += rowPoints

        winningNumbers = make(map[string]bool)
    }


    copies := make(map[int]int)
    sumCards := 0
    for i := 1; i <= len(gameMatches); i++ {
        matches := gameMatches[i]
        sumCards += copies[i] + 1  
        if(matches > 0) {
            for j := 1; j <= matches; j++ {
                copies[i + j] += copies[i] + 1
            }
        }
    }

    
    fmt.Println(sum)
    fmt.Println(sumCards)

    fmt.Println(time.Since(start))
}
    
