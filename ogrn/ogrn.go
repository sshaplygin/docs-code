package ogrn

import (
	"fmt"
)

// Validate check to valid OGRN format
// example: input format is 1027700132195
func Validate(ogrn string) (bool, error) {
	ogrnData, err := ParseOGRN(Legal, ogrn)
	if err != nil {
		return false, fmt.Errorf("parse %s model: %w", packageName, err)
	}

	return ogrnData.IsValid()
}

func Generate() string {
	return NewOGRN(Legal).String()
}
