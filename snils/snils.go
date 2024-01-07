package snils

import (
	"fmt"
)

// Validate check to valid SNILS format
// example: input format is 112-233-445 95
func Validate(snils string) (bool, error) {
	snilsData, err := ParseSNILS(snils)
	if err != nil {
		return false, fmt.Errorf("parse %s model: %w", packageName, err)
	}

	return snilsData.IsValid()
}

// Generate generate random
func Generate() string {
	return NewSNILS().String()
}
