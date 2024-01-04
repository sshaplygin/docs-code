package okato

import (
	"fmt"
)

// Validate check to valid OKATO format
// example: input format is 17 205 000 000
func Validate(okato string) (bool, error) {
	okatoData, err := ParseOKATO(okato)
	if err != nil {
		return false, fmt.Errorf("create %s model: %w", packageName, err)
	}

	return okatoData.IsValid()
}

// IsExist check to valid OKATO format and check to used code.
// Example valid format is 01 201 802 003 - Алтайский край/Районы Алтайского края/Алейский район/Сельсоветы Алейского р-на/п Мамонтовский
func IsExist(bik string) (bool, error) {
	okatoData, err := ParseOKATO(bik)
	if err != nil {
		return false, fmt.Errorf("create %s model: %w", packageName, err)
	}

	isValid, err := okatoData.IsValid()
	if err != nil {
		return false, fmt.Errorf("check valid %s model: %w", packageName, err)
	}

	if !isValid {
		return false, fmt.Errorf("invalid %s model", packageName)
	}

	return okatoData.IsExist()
}

func Generate() string {
	panic("not implemented!")
}
