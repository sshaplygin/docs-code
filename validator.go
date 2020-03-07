package go_codes_validator

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrInvalidINNLength = errors.New("invalid inn length")
	ErrInvalidBIKLength = errors.New("invalid bik length")

	ErrInvalidValue = errors.New("invalid code value")

	ErrInvalidBIKCountryCode = errors.New("invalid bik country code")
)

func strToArr(str string) ([]int, error) {
	numbers := strings.Split(str, "")
	arr := make([]int, 0, len(numbers))
	for _, number := range numbers {
		number, err := strconv.Atoi(number)
		if err != nil {
			return []int{}, ErrInvalidValue
		}
		arr = append(arr, number)
	}
	return arr, nil
}

func ValidateINN(inn string) (bool, error) {
	if len(inn) != 10 && len(inn) != 12 {
		return false, ErrInvalidINNLength
	}
	innArr, err := strToArr(inn)
	if err != nil {
		return false, err
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
	if len(bik) != 9 {
		return false, ErrInvalidBIKLength
	}
	bikArr, err := strToArr(bik)
	if err != nil {
		return false, err
	}
	if bikArr[0] != 0 || bikArr[1] != 4 {
		return false, ErrInvalidBIKCountryCode
	}
	// special code
	if bikArr[6] == 0 && bikArr[7] == 1 && bikArr[8] == 2 {
		return true, nil
	}
	latestTriadStr := bik[6:]
	code, _ := strconv.Atoi(latestTriadStr)
	return code >= 50 && code < 1000, nil
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

func ValidateKPP(kpp string) (bool, error) {
	return false, nil
}
