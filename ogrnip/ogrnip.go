package ogrnip

import (
	"strconv"

	ru_doc_code "github.com/mrfoe7/ru-doc-code"
)

// Validate check to valid OGRNIP format
// example: input format is 304500116000157
func Validate(ogrnip string) (bool, error) {
	if len(ogrnip) != 15 {
		pkg, err := ru_doc_code.GetPackageName()
		if err != nil {
			return false, err
		}
		return false, &ru_doc_code.CommonError{
			Method: pkg,
			Err:    ru_doc_code.ErrInvalidLength,
		}
	}

	ogrnipArr, err := ru_doc_code.StrToArr(ogrnip)
	if err != nil {
		return false, err
	}

	if ogrnipArr[0] != 3 && ogrnipArr[0] != 4 {
		return false, ru_doc_code.ErrInvalidValue
	}

	code, _ := strconv.Atoi(ogrnip[:len(ogrnip)-1])
	return ogrnipArr[len(ogrnip)-1] == code%13%10, nil
}
