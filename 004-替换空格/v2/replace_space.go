package main

import (
	"unicode/utf8"
	"strconv"
)

func ReplaceSpace(s string) (replaced string) {
	for _, Rune := range s {
		spaceRune, _ := utf8.DecodeRuneInString(" ")
		if Rune == spaceRune {
			Rune = spaceRune
		}

		replaced = replaced + strconv.QuoteRune(Rune)
	}
	return replaced
}
