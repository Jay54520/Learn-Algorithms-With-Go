package main

import "strings"

func ReplaceSpace(s string) (replaced string)  {
	return strings.Replace(s, " ", "%20", -1)
}
