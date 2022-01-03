package main

import (
	"fmt"
	"log"

	"github.com/FredySosa/advent-of-code-2021/day1"
	"github.com/FredySosa/advent-of-code-2021/day2"
	"github.com/FredySosa/advent-of-code-2021/day3"
)

func main() {
	//runDay1()
	//runDay2()
	runDay3()
}

func runDay3() {
	gama, epsilon, err := day3.GetGamaEpsilonNumbers("./day3/input.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Gama: %d. Epsilon: %d. Answer day 3.1: %d\n", gama, epsilon, (gama * epsilon))

	oxGen, coScrub, err := day3.GetOxGenCOScrub("./day3/input.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Oxigen Generator: %d. CO2 Scrubber: %d. Answer day 3.2: %d\n", oxGen, coScrub, (oxGen * coScrub))
}

func runDay2() {
	horizontal, depth, err := day2.GetHorizontalDepth("./day2/input.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Horizontal: %d. Depth: %d. Answer day 2.1: %d\n", horizontal, depth, (horizontal * depth))
	horizontal, depth, aim, err := day2.GetHorizontalDepthAim("./day2/input.txt")
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Horizontal: %d. Depth: %d. Aim: %d. Answer day 2.1: %d\n", horizontal, depth, aim, (horizontal * depth))
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
