package iteration

import (
	"strings"
)

func Repeat(character string, count int) string {
	var repeated strings.Builder
	for i := 0; i < count; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func CopyWordRepeat(word string, count int) string {
	var repeatedWord string
	for i := 0; i < count; i++ {
		repeatedWord = strings.Clone(word)
	}
	return repeatedWord
}
