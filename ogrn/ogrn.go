package ogrn

import (
	"strconv"

	ru_doc_code "github.com/sshaplygin/ru-doc-code"
)

// Validate check to valid OGRN format
// example: input format is 1027700132195
func Validate(ogrn string) (bool, error) {
	if len(ogrn) != 13 {
		name, err := ru_doc_code.GetModuleName()
		if err != nil {
			return false, err
		}
		return false, &ru_doc_code.CommonError{
			Method: name,
			Err:    ru_doc_code.ErrInvalidLength,
		}
	}

	ogrnArr, err := ru_doc_code.StrToArr(ogrn)
	if err != nil {
		return false, err
	}

	code, _ := strconv.Atoi(ogrn[:12])
	return ogrnArr[len(ogrn)-1] == code%11%10, nil
}
