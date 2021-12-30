package day1

func CountIncreasedNumbers(filePath string) (int, error) {
	integers, err := parseFileAsIntegers(filePath)
	if err != nil {
		return 0, err
	}
	var incresed int
	previous := integers[0]
	for _, actual := range integers[1:] {
		if previous < actual {
			incresed++
		}
		previous = actual
	}
	return incresed, nil
}
