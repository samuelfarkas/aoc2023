package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	"unicode"
)

// [-1, -1], [-1, 0], [-1, 1]
// all possible directions for every cell
// [0, -1], [0, 0], [0, 1]
// [1, -1], [1, 0], [1, 1]
var dx = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var dy = []int{-1, 0, 1, -1, 1, -1, 0, 1}

func checkAllDirections(i, j int, schema [][]rune) (bool, int, int) {
    isAdjacentToSymbol, isGearSymbolX, isGearSymbolY := false, -1, -1
    for dir := 0; dir < len(dx); dir++ {
        dirX, dirY := i+dx[dir], j+dy[dir]
        if dirX >= 0 && dirX < len(schema) && dirY >= 0 && dirY < len(schema[0]) {
                if !isNumberOrPeriod(schema[dirX][dirY]) {
                    isAdjacentToSymbol = true
                    if schema[dirX][dirY] == '*' {
                        isGearSymbolX = dirX
                        isGearSymbolY = dirY
                    }
                }
            }
        }
    return isAdjacentToSymbol, isGearSymbolX, isGearSymbolY 
}

func biDirectionalNumberSearch(i, j int, schema [][]rune, num *string) {
    var number string
    var x int

    for x = j; x >= 0 && unicode.IsDigit(schema[i][x]); x-- {}

    for x++; x < len(schema[i]) && unicode.IsDigit(schema[i][x]); x++ {
        number += string(schema[i][x])
    }

    fmt.Println("number", number, i, j)
}

func isNumberOrPeriod(chr rune) bool {
	return chr == '.' || unicode.IsDigit(chr)
}

func main() {
	start := time.Now()

	file, err := os.Open("test")

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

	isAdjacentToSymbol := false
    isGearSymbolX, isGearSymbolY := -1, -1 
	for i, row := range schema {
		num := ""
		for j, char := range row {
			if unicode.IsNumber(char) {
				num += string(char)
                isAdjacentToSymbol, isGearSymbolX, isGearSymbolY = checkAllDirections(i, j, schema)
			}

			if (!unicode.IsDigit(char) && num != "" && isAdjacentToSymbol) || (num != "" && isAdjacentToSymbol && j == len(row)-1) {
				number, err := strconv.Atoi(num)
				if err != nil {
					fmt.Println("error converting string to int")
				}
                if(isGearSymbolX != -1 && isGearSymbolY != -1) {
                    biDirectionalNumberSearch(isGearSymbolX, isGearSymbolY, schema, &num)
                }
				sum += number
				isAdjacentToSymbol = false
			}

			if !unicode.IsDigit(char) {
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
