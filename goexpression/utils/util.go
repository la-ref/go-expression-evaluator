package utils

type Char string

// Check if a character is a number
func IsCharANumber(value Char) bool {
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

// Check if a string or a character is a number
func ISANumber[T string | Char](value T) bool {
	for _, v := range value {
		if !IsCharANumber(Char(v)) {
			return false
		}
	}
	return true
}

// Check if a uint8 value is an ascii value
func IsCharANumberASCII(value uint8) bool {
	n9 := uint8('9')
	n0 := uint8('0')
	if value > n9 || value < n0 {
		return false
	}

	return true
}
