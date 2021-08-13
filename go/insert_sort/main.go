package main

import "log"

var cards = [6]int{5, 2, 4, 6, 1, 3}

func main() {
	for i := 1; i < len(cards); i++ {
		for j := i; j > 0 && cards[j-1] > cards[j]; j-- {
			cards[j], cards[j-1] = cards[j-1], cards[j]
		}
	}
	log.Println(cards)
}
