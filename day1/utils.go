package day1

import (
	"bufio"
	"os"
	"strconv"
)

func parseFileAsIntegers(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	toReturn := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		intStr := scanner.Text()
		parsed, err := strconv.Atoi(intStr)
		if err != nil {
			return nil, err
		}
		toReturn = append(toReturn, parsed)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return toReturn, nil
}
