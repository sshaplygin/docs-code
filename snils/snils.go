package snils

import (
	"strconv"
	"strings"

	ru_doc_code "github.com/mrfoe7/go-codes-validator"
)

// Validate check to valid SNILS format
// example: input format is 112-233-445 95
func Validate(snils string) (bool, error) {
	if len(snils) != 14 {
		return false, ru_doc_code.ErrInvalidSNILSLength
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
