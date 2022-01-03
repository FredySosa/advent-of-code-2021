package day3

import (
	"strconv"
)

func GetOxGenCOScrub(filePath string) (int, int, error) {
	binaryNumbers, err := parseFileAsStrings(filePath)
	if err != nil {
		return 0, 0, err
	}

	oxGenBinary := GetNumber(binaryNumbers, 0, 0)
	coScrubBinary := GetNumber(binaryNumbers, 0, 1)

	oxGen, err := strconv.ParseInt(oxGenBinary, 2, 64)
	if err != nil {
		return 0, 0, err
	}

	coScrub, err := strconv.ParseInt(coScrubBinary, 2, 64)
	if err != nil {
		return 0, 0, err
	}

	return int(oxGen), int(coScrub), nil
}

func GetNumber(numbersToSearch []string, number, bitCriteria int) string {
	if len(numbersToSearch) == 1 {
		return numbersToSearch[0]
	}

	counter := make(map[rune]int)
	for _, binary := range numbersToSearch {
		counter[rune(binary[number])]++
	}

	mostCommon := '1'
	if counter['0'] > counter['1'] {
		mostCommon = '0'
	}

	numbersToKeep := make([]string, 0)
	for _, binary := range numbersToSearch {
		isMostCommon := rune(binary[number]) == mostCommon

		if isMostCommon && bitCriteria == 0 {
			numbersToKeep = append(numbersToKeep, binary)
		}
		if !isMostCommon && bitCriteria == 1 {
			numbersToKeep = append(numbersToKeep, binary)
		}
	}

	return GetNumber(numbersToKeep, number+1, bitCriteria)
}
