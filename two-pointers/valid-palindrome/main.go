package main

import (
	"regexp"
	"strings"
)

func main() {

}

func isPalindrome(s string) bool {

	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	str := strings.ToLower(reg.ReplaceAllString(s, ""))

	if len(str) <= 0 {
		return true
	}

	j := len(str) - 1

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[j] {
			return false
		}
		j--
	}

	return true
}
