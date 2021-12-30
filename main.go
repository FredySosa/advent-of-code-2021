package main

import (
	"fmt"
	"log"

	"github.com/FredySosa/advent-of-code-2021/day1"
)

func main() {
	runDay1()
}

func runDay1() {
	totalIncreased, err := day1.CountIncreasedNumbers("./day1/input.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Day 1.1 answer:", totalIncreased)
	totalTripletsIncreased, err := day1.CountIncreasedTriplets("./day1/input.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Day 1.2 answer:", totalTripletsIncreased)
}
