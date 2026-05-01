package vo

import (
	errs "dddcarpark/internal/errors"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Description string

func NewDescription(value string) (Description, error) {
	if utf8.RuneCountInString(value) < 2 || utf8.RuneCountInString(value) > 300 {
		return "", errs.ErrInvalidDescriptionLength
	}

	r, size := utf8.DecodeRuneInString(value)

	description := Description(string(unicode.ToUpper(r)) + strings.ToLower(value[size:]))

	return description, nil
}
