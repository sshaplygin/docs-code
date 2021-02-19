package kpp

import ru_doc_code "github.com/mrfoe7/ru-doc-code"

// Validate check to valid KPP format
// example: input format is 773643301
func Validate(kpp string) (bool, error) {
	if len(kpp) != 9 {
		return false, ru_doc_code.ErrInvalidKPPLength
	}

	_, err := ru_doc_code.StrToArr(kpp)
	if err != nil {
		return false, err
	}

	return true, nil
}
