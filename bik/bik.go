package bik

import (
	"fmt"
)

// Validate check to valid BIK format.
// Example valid format is 044525225
func Validate(bik string) (bool, error) {
	bikData, err := ParseBIK(bik)
	if err != nil {
		return false, fmt.Errorf("parse %s model: %w", packageName, err)
	}

	return bikData.IsValid()
}

// Exists check to valid BIK format and check to used code.
// Example valid format is 044525677 - АО "Яндекс Банк".
func Exists(bik string) (bool, error) {
	_, ok := existsBIKs[bik]
	if ok {
		return true, nil
	}

	bikData, err := ParseBIK(bik)
	if err != nil {
		return false, fmt.Errorf("parse %s model: %w", packageName, err)
	}

	isValid, err := bikData.IsValid()
	if err != nil {
		return false, fmt.Errorf("check valid %s model: %w", packageName, err)
	}

	if !isValid {
		return false, fmt.Errorf("invalid %s model", packageName)
	}

	return bikData.Exists()
}

// Generate method generate a valid BIK code, but possible usaged or not usaged in reality.
// Method guaranteed that code will be valid, but not guaranteed that code will be exists.
func Generate() string {
	return NewBIK().String()
}
