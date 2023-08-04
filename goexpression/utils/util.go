package utils

func IsCharANumber(value string) bool {
	if len(value) <= 0 || len(value) > 1 {
		return false
	}
	ascii := int(value[0])
	n9 := int('9')
	n0 := int('0')
	if ascii > n9 || ascii < n0 {
		return false
	}

	return true
}
