package day2

func GetHorizontalDepthAim(filePath string) (int, int, int, error) {
	movements, err := parseFileToDataStruct(filePath)
	if err != nil {
		return 0, 0, 0, err
	}
	var horizontal, depth, aim int

	for _, movement := range movements {
		switch movement.Move {
		case "forward":
			horizontal += movement.Amount
			depth += aim * movement.Amount
		case "up":
			aim -= movement.Amount
		case "down":
			aim += movement.Amount
		}
	}

	return horizontal, depth, aim, nil
}
