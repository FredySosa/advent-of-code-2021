package day1

func CountIncreasedTriplets(filePath string) (int, error) {
	integers, err := parseFileAsIntegers(filePath)
	if err != nil {
		return 0, err
	}

	var incresed int
	previous := integers[0] + integers[1] + integers[2]
	for i := 1; i < len(integers)-2; i++ {
		actual := integers[i] + integers[i+1] + integers[i+2]
		if actual > previous {
			incresed++
		}
		previous = actual
	}

	return incresed, nil
}
