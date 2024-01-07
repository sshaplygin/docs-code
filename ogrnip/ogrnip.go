package ogrnip

import (
	"fmt"

	"github.com/sshaplygin/docs-code/ogrn"
)

// Validate check to valid OGRNIP format
// example: input format is 304500116000157
func Validate(ogrnip string) (bool, error) {
	ogrnData, err := ogrn.ParseOGRN(ogrn.Physical, ogrnip)
	if err != nil {
		return false, fmt.Errorf("parse %s model: %w", packageName, err)
	}

	return ogrnData.IsValid()
}

func Generate() string {
	return ogrn.NewOGRN(ogrn.Physical).String()
}
