package convert

import "strconv"

func StringToInt(str string) int {
	if len(str) == 0 {
		return 0
	}

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	if i <= 0 {
		return 0
	}

	return i
}

func StringToFloat64(str string) float64 {
	if len(str) == 0 {
		return 0
	}

	i, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}

	return i
}

func StringToBool(str string) bool {
	if len(str) == 0 {
		return false
	}

	b, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}

	return b
}
