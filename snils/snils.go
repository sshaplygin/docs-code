package snils

import (
	"strconv"
	"strings"

	"github.com/sshaplygin/ru-doc-code/models"
	"github.com/sshaplygin/ru-doc-code/utils"
)

// Validate check to valid SNILS format
// example: input format is 112-233-445 95
func Validate(snils string) (bool, error) {
	if len(snils) != 14 {
		return false, &models.CommonError{
			Method: packageName,
			Err:    models.ErrInvalidLength,
		}
	}

	fSnils := strings.ReplaceAll(snils, "-", "")
	fSnils = strings.ReplaceAll(fSnils, " ", "")

	if len(fSnils) != 11 {
		return false, ErrInvalidFormattedLength
	}

	snilsArr, err := utils.StrToArr(fSnils)
	if err != nil {
		return false, err
	}

	hashSum := 0
	hashLen := len(fSnils) - 2
	code, _ := strconv.Atoi(fSnils[hashLen:])
	for i, v := range snilsArr[:hashLen] {
		hashSum += v * (hashLen - i)
	}

	return hashSum%101 == code, nil
}

// Generate generate random
func Generate() string {
	panic("not implemented!")
}
