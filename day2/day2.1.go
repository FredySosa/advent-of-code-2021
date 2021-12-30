package day2

func GetHorizontalDepth(filePath string) (int, int, error) {
	movements, err := parseFileToDataStruct(filePath)
	if err != nil {
		return 0, 0, err
	}
	var horizontal int
	var depth int

	for _, move := range movements {
		switch move.Move {
		case "forward":
			horizontal += move.Amount
		case "up":
			depth -= move.Amount
		case "down":
			depth += move.Amount
		}
	}

	return horizontal, depth, nil
}
