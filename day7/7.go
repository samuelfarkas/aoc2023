package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)   

var cards_pt1 = []byte{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'} 
var cards_pt2 = []byte{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'} 

var cards = cards_pt2

var (
	HighCard = []int{1, 1, 1, 1, 1}
	OnePair = []int{2, 1, 1, 1}
	TwoPair = []int{2, 2, 1}
	ThreeOfAKind = []int{3, 1, 1}
	FullHouse = []int{3, 2}
	FourOfAKind = []int{4, 1}
	FiveOfAKind = []int{5}
)

var CardPatterns = [][]int{
    HighCard,
    OnePair,
    TwoPair,
    ThreeOfAKind,
    FullHouse,
    FourOfAKind,
    FiveOfAKind,
}

type Hand struct {
    cardsIdx [5]int
    bid int
    patternIdx int
}


func makeHandPattern(hand string, applyJoker bool) []int {
    m := make(map[int]int)
    
    for _, card := range hand {
        cardIdx := bytes.IndexByte(cards, byte(card))
        m[cardIdx] += 1
    }

    if m[12] == 5 && applyJoker {
        return FiveOfAKind
    }

    pattern := make([]int, 0) 
    jokerCount := 0
    for key, v := range m {
        if(applyJoker && key == 12) {
            jokerCount = v
            continue
        }
        pattern = append(pattern, v)
    }
    sort.Slice(pattern, func(i, j int) bool {
        return pattern[i] > pattern[j]
    })

    if(applyJoker) {
        pattern[0] += jokerCount
    }
    return pattern
}

func identifyPatternIdx(patternOnHand []int) int {
    for patternIdx, pattern := range CardPatterns {
        for j, num := range patternOnHand { 
            if num != pattern[j] {
                break
            }             
            if j == len(pattern) - 1 {
                return patternIdx  
            }
        }
    }
    return -1
}

func applyOrdering(hands *[]*Hand) {
    h := *hands
    sort.SliceStable(h, func(i, j int) bool {
        current := *h[i]
        next := *h[j]
        // if same type, then rank by card indices
        if current.patternIdx == next.patternIdx {
            for k := range current.cardsIdx {
                if current.cardsIdx[k] != next.cardsIdx[k] {
                    // highest card first
                    return current.cardsIdx[k] > next.cardsIdx[k]
                }
            }
        }
        // else rank by pattern - highest pattern first
        return current.patternIdx < next.patternIdx 
    })
}


func main() {
    start := time.Now()
	file, err := os.Open("input")

	if err != nil {
		fmt.Println("error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

    hands := make([]*Hand, 0)
    hands2 := make([]*Hand, 0)

    for scanner.Scan() {
        line := scanner.Text()
        splitBySpace := strings.Split(line, " ")
        handString, bid := splitBySpace[0], splitBySpace[1] 

        hand := &Hand{[5]int{}, 0, 0}
        hand2 := &Hand{[5]int{}, 0, 0}
        
        // set bid
        bidNum, err := strconv.Atoi(bid)
        if err != nil {
            fmt.Println("error converting bid to int")
        }

        for i, card := range handString {
            hand.cardsIdx[i] = bytes.IndexByte(cards, byte(card))
            hand2.cardsIdx[i] = bytes.IndexByte(cards, byte(card)) 
        }

        hand.bid = bidNum
        hand2.bid = bidNum
        handPattern := makeHandPattern(handString, false)
        handPattern_pt2 := makeHandPattern(handString, true)

        hand.patternIdx = identifyPatternIdx(handPattern)
        hand2.patternIdx = identifyPatternIdx(handPattern_pt2)

        hands = append(hands, hand)
        hands2 = append(hands2, hand2)
    }


    applyOrdering(&hands) 
    applyOrdering(&hands2)

    winning := 0
    winning_pt2 := 0
    for i, hand := range hands {
        winning += hand.bid * (i + 1)
        winning_pt2 += hands2[i].bid * (i + 1)
    }

    fmt.Println("Winning:", winning)
    fmt.Println("Winning pt2:", winning_pt2)


    fmt.Println("Time elapsed: ", time.Since(start))

}
