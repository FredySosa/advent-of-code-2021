package day3

import (
	"bufio"
	"os"
)

func parseFileAsStrings(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	toReturn := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		toReturn = append(toReturn, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return toReturn, nil
}
