package valueobject

import (
	"dddcarpark/internal/errors"
	"strings"
	"unicode/utf8"
)

type VinNumber string

func NewVinNumber(vin string) (VinNumber, error) {
	if utf8.RuneCountInString(vin) != 17 {
		return "", errors.ErrInvalidVinNumberLength
	}
	if strings.Contains(vin, "I") || strings.Contains(vin, "O") || strings.Contains(vin, "Q") {
		return "", errors.ErrInvalidVinNumberCharacters
	}
	return VinNumber(vin), nil
}
