package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"unicode"
)

type TrieNode struct {
	children     map[rune]*TrieNode
	isWord       bool
	numericValue int
}

func createNode() *TrieNode {
    // Create trie node, where children is a map of runes (keys) to trie nodes pointers (values)
    // return pointer to trie node
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

// this syntax is called a method receiver and it allow calling this function on a TrieNode object
func (t *TrieNode) insertWord(word string, numericValue int) {
	node := t
	for _, ch := range word {
		if node.children[ch] == nil {
			node.children[ch] = createNode()
		}
		node = node.children[ch]
		node.numericValue = -1
	}
	node.numericValue = numericValue
	node.isWord = true
}

func (t *TrieNode) search(word string) (bool, int) {
	node := t
	for _, ch := range word {
		node = node.children[ch]
		if node == nil {
			return false, -1
		}
	}
	return node.isWord, node.numericValue
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
	digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	trieRoot := createNode()

	for i, digit := range digits {
		trieRoot.insertWord(digit, i+1)
	}

	lineIdx := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineIdx++
		var firstNum rune
		var lastNum rune

		for i, char := range line {
			if unicode.IsDigit(char) {
				if firstNum == 0 {
					firstNum = char
				}

				lastNum = char
			}

			var digit string
			for j := i; j < len(line); j++ {
				digit += string(line[j])
				isWord, numericValue := trieRoot.search(digit)
				if isWord {
					if firstNum == 0 {
						firstNum = rune(numericValue + 48)
					}

					lastNum = rune(numericValue + 48)
				}
			}
		}
		sum += int(firstNum-'0')*10 + int(lastNum-'0')
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error scanning file")
	}

	fmt.Println("Sum:", sum)
	fmt.Println("Running time", time.Since(start))
}
