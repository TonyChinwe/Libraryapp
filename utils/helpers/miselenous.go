package helpers

import (
	"errors"
	"strings"
	"unicode"
)

func IsEmpty(data []interface{}) bool {
	return len(data) == 0
}

func NewError(message string) error {
	newString := strings.ToTitleSpecial(unicode.SpecialCase{}, message)
	return errors.New(newString)
}
