package kpp

import "fmt"

// Validate check to valid KPP format
// example: input format is 773643301
func Validate(kpp string) (bool, error) {
	kppData, err := ParseKPP(kpp)
	if err != nil {
		return false, fmt.Errorf("parse %s model: %w", packageName, err)
	}

	return kppData.IsValid()
}

func Generate() string {
	return NewKPP().String()
}
