package go_codes_validator

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errInvalidLength = errors.New("invalid inn length")
	errInvalidValue  = errors.New("invalid inn value")
)

func ValidateINN(inn string) (bool, error) {
	if len(inn) != 10 && len(inn) != 12 {
		return false, errInvalidLength
	}
	innNumbers := strings.Split(inn, "")
	innArr := make([]int, 0, len(inn))
	for _, str := range innNumbers {
		number, err := strconv.Atoi(str)
		if err != nil {
			return false, errInvalidValue
		}
		innArr = append(innArr, number)
	}
	if len(inn) == 10 {
		controlNumber := ((2*innArr[0] + 4*innArr[1] + 10*innArr[2] + 3*innArr[3] + 5*innArr[4] + 9*innArr[5] + 4*innArr[6] + 6*innArr[7] + 8*innArr[8]) % 11) % 10
		return controlNumber == innArr[len(innArr)-1], nil
	}
	firstControlNumber := ((7*innArr[0] + 2*innArr[1] + 4*innArr[2] + 10*innArr[3] + 3*innArr[4] + 5*innArr[5] + 9*innArr[6] + 4*innArr[7] + 6*innArr[8] + 8*innArr[9]) % 11) % 10
	secondControlNumber := ((3*innArr[0] + 7*innArr[1] + 2*innArr[2] + 4*innArr[3] + 10*innArr[4] + 3*innArr[5] + 5*innArr[6] + 9*innArr[7] + 4*innArr[8] + 6*innArr[9] + 8*innArr[10]) % 11) % 10
	return firstControlNumber == innArr[len(innArr)-2] && secondControlNumber == innArr[len(innArr)-1], nil
}

func ValidateBIK(bik string) (bool, error) {
	return false, nil
}

func ValidateOGRN(ogrn string) (bool, error) {
	return false, nil
}

func ValidateOGRNIP(ogrnip string) (bool, error) {
	return false, nil
}

func ValidateSNILS(snils string) (bool, error) {
	return false, nil
}
