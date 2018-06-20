package main

func ReplaceSpace(s string) (replaced string) {
	replacedRune := []rune("")
	repSlice := []rune("%20")
	for _, runeValue := range s {
		if string(runeValue) == " " {
			replacedRune = append(replacedRune, repSlice...)
		} else {
			replacedRune = append(replacedRune, runeValue)
		}
	}
	replaced = string(replacedRune)
	return
}
