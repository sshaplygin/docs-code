package inn

import (
	"fmt"

	"github.com/sshaplygin/docs-code/utils"
)

// Validate check to valid inn from input string.
// example: input format is 7707083893
func Validate(inn string) (bool, error) {
	innData, err := ParseINN(inn)
	if err != nil {
		return false, fmt.Errorf("parse %s model: %w", packageName, err)
	}

	return innData.IsValid()
}

// Generate generate random type inn string value
func Generate() string {
	return NewINN(INNType(utils.RandomDigits(1) % 3)).String()
}

// GenerateLegal generate legal type inn string value
func GenerateLegal() string {
	return NewINN(INNType(utils.RandomDigits(1)%2 + 1)).String()
}

// GeneratePhysical generate physical type inn string value
func GeneratePhysical() string {
	return NewINN(INNType(Physical)).String()
}
