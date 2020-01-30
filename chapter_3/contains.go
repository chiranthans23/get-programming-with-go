package chapter_3

import (
	"strings"
)

//ContainsString checks if a given string is substring of another string
func ContainsString(BigString string, ele string) bool {
	return strings.Contains(BigString, ele)
}
