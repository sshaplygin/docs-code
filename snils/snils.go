package snils

import (
	"strconv"
	"strings"

	ru_doc_code "github.com/mrfoe7/ru-doc-code"
)

// Validate check to valid SNILS format
// example: input format is 112-233-445 95
func Validate(snils string) (bool, error) {
	if len(snils) != 14 {
		name, err := ru_doc_code.GetModuleName()
		if err != nil {
			return false, err
		}
		return false, &ru_doc_code.CommonError{
			Method: name,
			Err:    ru_doc_code.ErrInvalidLength,
		}
	}

	fSnils := strings.ReplaceAll(snils, "-", "")
	fSnils = strings.ReplaceAll(fSnils, " ", "")

	if len(fSnils) != 11 {
		return false, ru_doc_code.ErrInvalidFormattedSNILSLength
	}

	snilsArr, err := ru_doc_code.StrToArr(fSnils)
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
	return ""
}
