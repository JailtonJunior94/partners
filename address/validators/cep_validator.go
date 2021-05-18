package validators

import (
	"regexp"
	"strings"
)

func ValidateInputLength(cepRaw string) (status bool) {
	const cepSize = 8
	cepLength := len(cepRaw)
	if cepLength <= cepSize {
		status = true
	} else {
		status = false
	}
	return status
}

func RemoveSpecialCharacters(cepRaw string) (cepParsed string) {
	rule := regexp.MustCompile(`\D+`)
	cepParsed = rule.ReplaceAllString(cepRaw, "")
	return cepParsed
}

func LeftPadWithZeros(cepRaw string) (cepParsed string) {
	const cepSize = 8
	cepLength := len(cepRaw)
	timesToRepeat := cepSize - cepLength
	pad := strings.Repeat("0", timesToRepeat)
	cepParsed = pad + cepRaw
	return cepParsed
}
