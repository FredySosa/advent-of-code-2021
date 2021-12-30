package day2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Move   string
	Amount int
}

func parseFileToDataStruct(fileName string) ([]Data, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	toReturn := make([]Data, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dataStr := scanner.Text()
		arrayOfData := strings.Fields(dataStr)
		amountMove, err := strconv.Atoi(arrayOfData[1])
		if err != nil {
			return nil, err
		}
		toReturn = append(toReturn, Data{
			Move:   arrayOfData[0],
			Amount: amountMove,
		})
	}
	return toReturn, nil
}
