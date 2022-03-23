package util

func IsElementPresent(element int, arr []int) bool {
	for _, i := range arr {
		if i == element {
			return true
		}
	}
	return false
}
