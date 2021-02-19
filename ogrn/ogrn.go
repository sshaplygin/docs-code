package ogrn

import (
	"strconv"

	ru_doc_code "github.com/mrfoe7/ru-doc-code"
)

// Validate check to valid OGRN format
// example: input format is 1027700132195
func Validate(ogrn string) (bool, error) {
	if len(ogrn) != 13 {
		return false, ru_doc_code.ErrInvalidOGRNLength
	}

	ogrnArr, err := ru_doc_code.StrToArr(ogrn)
	if err != nil {
		return false, err
	}

	code, _ := strconv.Atoi(ogrn[:12])
	return ogrnArr[len(ogrn)-1] == code%11%10, nil
}
