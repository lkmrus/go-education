package utils

func ConvertStringToUint(s string) uint {
	var result uint
	for _, c := range s {
		result = result*10 + uint(c-'0')
	}
	return result
}
