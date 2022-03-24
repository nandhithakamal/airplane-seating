package util

import "log"

func IsElementPresent(element int, arr []int) bool {
	for _, i := range arr {
		if i == element {
			return true
		}
	}
	return false
}

func HandleError(err error) {
	log.Fatal(err)
}
