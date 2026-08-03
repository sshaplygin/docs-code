package ogrn

import (
	"fmt"

	"github.com/sshaplygin/docs-code/models"
)

// Validate check to valid OGRN format. It validates 13-digit OGRN codes of both
// legal entities and government bodies. Individual entrepreneur codes (OGRNIP)
// have their own ogrnip package.
// example: input format is 1027700132195
func Validate(ogrn string) (bool, error) {
	ogrnType, err := detectType(ogrn)
	if err != nil {
		return false, err
	}

	ogrnData, err := ParseOGRN(ogrnType, ogrn)
	if err != nil {
		return false, fmt.Errorf("parse %s model: %w", packageName, err)
	}

	return ogrnData.IsValid()
}

// detectType infers the OGRN type from the leading code digit. Physical
// (individual entrepreneur) codes are rejected: they belong to the ogrnip package.
func detectType(ogrn string) (OGRNType, error) {
	if len(ogrn) == 0 {
		return 0, &models.CommonError{Method: packageName, Err: models.ErrInvalidLength}
	}

	first := ogrn[0]
	if first < '0' || first > '9' {
		return 0, &models.CommonError{Method: packageName, Err: models.ErrInvalidValue}
	}

	ogrnType, ok := supportedCodes[CodeType(first-'0')]
	if !ok || ogrnType == Physical {
		return 0, ErrInvalidCodeType
	}

	return ogrnType, nil
}

func Generate() string {
	return NewOGRN(Legal).String()
}
