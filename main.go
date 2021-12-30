package main

import (
	"fmt"
	"log"

	"github.com/FredySosa/advent-of-code-2021/day1"
	"github.com/FredySosa/advent-of-code-2021/day2"
)

func main() {
	//runDay1()
	runDay2()
}

func runDay2() {
	horizontal, depth, err := day2.GetHorizontalDepth("./day2/input.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Horizontal: %d. Depth: %d. Answer day 2.1: %d\n", horizontal, depth, (horizontal * depth))
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
