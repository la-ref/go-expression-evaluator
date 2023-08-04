package utils

func IsCharANumber(value string) bool {
	if len(value) <= 0 || len(value) > 1 {
		return false
	}
	ascii := uint8(value[0])
	n9 := uint8('9')
	n0 := uint8('0')
	if ascii > n9 || ascii < n0 {
		return false
	}

	return true
}

func IsCharANumberASCII(value uint8) bool {
	n9 := uint8('9')
	n0 := uint8('0')
	if value > n9 || value < n0 {
		return false
	}

	return true
}
