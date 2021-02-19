package inn

import (
	"strconv"
	"strings"

	ru_doc_code "github.com/mrfoe7/go-codes-validator"
)

const (
	lengthLegal    = 10
	lengthPhysical = 12
)

// Validate check to valid inn from input string.
// example: input format is 7707083893
func Validate(inn string) (bool, error) {
	if len(inn) != lengthLegal && len(inn) != lengthPhysical {
		return false, ru_doc_code.ErrInvalidINNLength
	}

	innArr, err := transformInn(inn)
	if err != nil {
		return false, err
	}
	if len(inn) == 10 {
		return hash10(innArr) == innArr[len(innArr)-1], nil
	}

	return hash11(innArr) == innArr[len(innArr)-2] && hash12(innArr) == innArr[len(innArr)-1], nil
}

// GenerateLegal generate legal type inn string value.
func GenerateLegal() string {
	inn := strconv.FormatInt(ru_doc_code.RandomDigits(9), 10)
	innArr, _ := transformInn(inn)

	return inn + strconv.Itoa(hash10(innArr))
}

// GeneratePhysical generate physical type inn string value.
func GeneratePhysical() string {
	inn := strconv.FormatInt(ru_doc_code.RandomDigits(10), 10)
	innArr, _ := transformInn(inn)

	hash1Num := hash11(innArr)
	innArr = append(innArr, hash1Num)

	return strings.Join([]string{
		inn,
		strconv.Itoa(hash1Num),
		strconv.Itoa(hash12(innArr)),
	}, "")
}

func transformInn(inn string) ([]int, error) {
	innNumbers := strings.Split(inn, "")
	innArr := make([]int, 0, len(inn))

	for _, str := range innNumbers {
		number, err := strconv.Atoi(str)
		if err != nil {
			return nil, ru_doc_code.ErrInvalidValue
		}
		innArr = append(innArr, number)
	}

	return innArr, nil
}

// Generate generate random type inn string value.
func Generate() string {
	if ru_doc_code.RandomDigits(1)%2 == 1 {
		return GeneratePhysical()
	}

	return GenerateLegal()
}

// IsEntrepreneur check inn for Entrepreneur type
func IsEntrepreneur(inn string) bool {
	return len(inn) == lengthPhysical
}

// IsCompany check inn for Company type
func IsCompany(inn string) bool {
	return len(inn) == lengthLegal
}

func hash10(innArr []int) int {
	return ((2*innArr[0] + 4*innArr[1] + 10*innArr[2] + 3*innArr[3] +
		5*innArr[4] + 9*innArr[5] + 4*innArr[6] + 6*innArr[7] + 8*innArr[8]) % 11) % 10
}

func hash11(innArr []int) int {
	return ((7*innArr[0] + 2*innArr[1] + 4*innArr[2] + 10*innArr[3] + 3*innArr[4] +
		5*innArr[5] + 9*innArr[6] + 4*innArr[7] + 6*innArr[8] + 8*innArr[9]) % 11) % 10
}

func hash12(innArr []int) int {
	return ((3*innArr[0] + 7*innArr[1] + 2*innArr[2] + 4*innArr[3] +
		10*innArr[4] + 3*innArr[5] + 5*innArr[6] + 9*innArr[7] + 4*innArr[8] +
		6*innArr[9] + 8*innArr[10]) % 11) % 10
}
