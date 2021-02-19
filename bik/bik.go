package bik

import (
	"strconv"

	ru_doc_code "github.com/mrfoe7/go-codes-validator"
)

// Validate check to valid BIK format
// example valid format is 044525225
func Validate(bik string) (bool, error) {
	if len(bik) != 9 {
		return false, ru_doc_code.ErrInvalidBIKLength
	}

	bikArr, err := ru_doc_code.StrToArr(bik)
	if err != nil {
		return false, err
	}

	if bikArr[0] != 0 || bikArr[1] != 4 {
		return false, ru_doc_code.ErrInvalidBIKCountryCode
	}

	// special code
	if bikArr[6] == 0 && bikArr[7] == 1 && bikArr[8] == 2 {
		return true, nil
	}

	latestTriadStr := bik[6:]
	code, _ := strconv.Atoi(latestTriadStr)

	return code >= 50 && code < 1000, nil
}
