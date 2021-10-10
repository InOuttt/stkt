package refactor

import (
	"strings"
)

func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			wordsAfterFirstBracket := str[indexFirstBracketFound:]
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				return wordsAfterFirstBracket[1:indexClosingBracketFound]
			}
		}
	}
	return ""
}
