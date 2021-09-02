package cmd

import (
	"regexp"
	"unicode"
)

// check if the word is Chinese
func isChinese(word string) bool {
	for _, val := range word {
		if unicode.Is(unicode.Han, val) {
			return true
		}
	}
	return false
}

// delete `\n`
func del(str string) string {
	r := regexp.MustCompile("\n")
	res := r.ReplaceAllString(str, "")
	return res
}
