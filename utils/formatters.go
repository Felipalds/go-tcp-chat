package utils

func IsRequestValid(buff []string) bool {
	if len(buff) <= 1 {
		return false
	}
	if len(buff[0]) <= 1 {
		return false
	}
	return true
}
