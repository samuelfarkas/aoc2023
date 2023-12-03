package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	"unicode"
)

func isNumberOrPeriod(chr rune) bool {
    return chr == '.' || unicode.IsDigit(chr)
}

func main() {
    start := time.Now()

    file, err := os.Open("input")

    if err != nil {
        fmt.Println("error opening file")
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var schema [][]rune
    for scanner.Scan() {
        line := scanner.Text()
        chars := []rune(line) 
        schema = append(schema, chars)
    }
    sum := 0

    // all possible directions for every cell
    // [-1, -1], [-1, 0], [-1, 1] 
    // [0, -1], [0, 0], [0, 1] 
    // [1, -1], [1, 0], [1, 1]
    dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
    dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}
    fmt.Println(len(schema), len(schema[0]))
        
    isAdjacentToSymbol := false
    for i, row := range schema {
        num := ""
        for j, char := range row {
            if unicode.IsNumber(char) {
                num += string(char)
                for dir := 0; dir < len(dx); dir++ {
                    dirX, dirY := i + dx[dir], j + dy[dir]
                    if dirX >= 0 && dirX < len(schema) && dirY >= 0 && dirY < len(schema[0]) {
                        if !isNumberOrPeriod(schema[dirX][dirY]) {
                            isAdjacentToSymbol = true
                        }
                    }
                }
            } 

            if (!unicode.IsDigit(char) && num != "" && isAdjacentToSymbol) || (num != "" && isAdjacentToSymbol && j == len(row) - 1) {
                number, err := strconv.Atoi(num)
                if err != nil {
                    fmt.Println("error converting string to int")
                }
                sum += number
                isAdjacentToSymbol = false
            } 

            if(!unicode.IsDigit(char)) {
                num = ""
            }
        }
    }



    fmt.Println(sum)

    if err := scanner.Err(); err != nil {
        fmt.Println("error scanning file")
    }



    fmt.Println(time.Since(start))
}
