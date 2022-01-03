package day3

import (
	"strconv"
)

func GetGamaEpsilonNumbers(filePath string) (int, int, error) {
	binaryNumbers, err := parseFileAsStrings(filePath)
	if err != nil {
		return 0, 0, err
	}
	counter := make(map[int]map[int]int, 0)
	for i := 0; i < len(binaryNumbers[0]); i++ {
		counter[i] = make(map[int]int)
	}

	for _, number := range binaryNumbers {
		for index, value := range number {
			toRepeat := value - 48
			counter[index][int(toRepeat)]++
		}
	}

	var gamaBinary, epsilonBinary string
	for i := 0; i < len(binaryNumbers[0]); i++ {
		nextGama := "0"
		nextEpsilon := "1"

		if counter[i][1] > counter[i][0] {
			nextGama, nextEpsilon = nextEpsilon, nextGama
		}

		gamaBinary += nextGama
		epsilonBinary += nextEpsilon
	}

	gama, err := strconv.ParseInt(gamaBinary, 2, 64)
	if err != nil {
		return 0, 0, err
	}

	epsilon, err := strconv.ParseInt(epsilonBinary, 2, 64)
	if err != nil {
		return 0, 0, err
	}

	return int(gama), int(epsilon), nil
}
