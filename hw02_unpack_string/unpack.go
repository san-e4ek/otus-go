package hw02unpackstring

import (
	"errors"
	"strconv"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	runes := []rune{}

	for index, char := range str {
		var nextChar rune

		if index+1 < len(str) {
			nextChar = rune(str[index+1])
		}

		switch {
		case unicode.IsDigit(char) && unicode.IsDigit(nextChar):
			return "", ErrInvalidString
		case unicode.IsDigit(char) && len(runes) == 0:
			return "", ErrInvalidString
		case unicode.IsSymbol(char) || unicode.IsLetter(char):
			runes = append(runes, char)
		case unicode.IsDigit(char):
			num, _ := strconv.Atoi(string(char))

			if num == 0 {
				runes = runes[:len(runes)-1]
			} else {
				for i := 0; i < num-1; i++ {
					runes = append(runes, runes[len(runes)-1])
				}
			}
		}
	}

	return string(runes), nil
}
