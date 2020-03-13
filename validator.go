package go_codes_validator

import (
	"strconv"
	"strings"
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

// IsINNValid
// example valid format is
func IsINNValid(inn string) (bool, error) {
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

// IsBIKValid
// example valid format is
func IsBIKValid(bik string) (bool, error) {
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

// IsOGRNValid
// example valid format is
func IsOGRNValid(ogrn string) (bool, error) {
	if len(ogrn) != 13 {
		return false, ErrInvalidOGRNLength
	}
	ogrnArr, err := strToArr(ogrn)
	if err != nil {
		return false, err
	}
	code, _ := strconv.Atoi(ogrn[:12])
	return ogrnArr[len(ogrn)-1] == code%11%10, nil
}

// IsOGRNIPValid check to valid OGRNIP format
// example valid format is
func IsOGRNIPValid(ogrnip string) (bool, error) {
	if len(ogrnip) != 15 {
		return false, ErrInvalidOGRNIPLength
	}
	ogrnipArr, err := strToArr(ogrnip)
	if err != nil {
		return false, err
	}
	if ogrnipArr[0] != 3 && ogrnipArr[0] != 4 {
		return false, ErrInvalidValue
	}
	code, _ := strconv.Atoi(ogrnip[:12])
	return ogrnipArr[len(ogrnip)-1] == code%13%10, nil
}

// IsSNILSValid check
// example valid format is `112-233-445 95`
func IsSNILSValid(snils string) (bool, error) {
	if len(snils) != 14 {
		return false, ErrInvalidSNILSLength
	}
	fSnils := strings.ReplaceAll(snils, "-", "")
	fSnils = strings.ReplaceAll(fSnils, " ", "")
	if len(fSnils) != 11 {
		return false, ErrInvalidFormattedSNILSLength
	}
	snilsArr, err := strToArr(fSnils)
	if err != nil {
		return false, err
	}
	hashSum := 0
	hashLen := len(snilsArr) - 2
	code, _ := strconv.Atoi(snils[hashLen:])
	for i, v := range snilsArr[:hashLen] {
		hashSum += v * (hashLen - i)
	}
	return hashSum%101 == code, nil
}

// IsKPPValid
// example valid format is
func IsKPPValid(kpp string) (bool, error) {
	if len(kpp) == 9 {
		return true, nil
	}
	return false, ErrInvalidKPPLength
}
