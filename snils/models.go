package snils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sshaplygin/docs-code/models"
	"github.com/sshaplygin/docs-code/utils"
)

const packageName = "snils"

type Numbers []int

const checkSumLength = 2

type checkSum int

type SNILSStruct struct {
	numbers  Numbers
	checkSum checkSum
}

func (ss *SNILSStruct) String() string {
	var res strings.Builder
	res.Grow(snilsFullLength)

	for i := 0; i < len(ss.numbers); i++ {
		if i%3 == 0 && i != 0 {
			res.WriteString("-")
		}
		res.WriteString(strconv.Itoa(ss.numbers[i]))
	}

	res.WriteString(" ")
	res.WriteString(utils.StrCode(int(ss.checkSum), checkSumLength))

	return res.String()
}

func (ss *SNILSStruct) IsValid() (bool, error) {
	return ss.calculateCheckSum() == ss.checkSum, nil
}

const (
	snilsFullLength   = 14
	snilsShrinkLength = 11
)

func ParseSNILS(snils string) (*SNILSStruct, error) {
	if len(snils) != snilsFullLength {
		return nil, &models.CommonError{
			Method: packageName,
			Err:    models.ErrInvalidLength,
		}
	}

	fSnils := strings.ReplaceAll(snils, "-", "")
	fSnils = strings.ReplaceAll(fSnils, " ", "")

	if len(fSnils) != snilsShrinkLength {
		return nil, ErrInvalidFormattedLength
	}

	snilsArr, err := utils.StrToArr(fSnils)
	if err != nil {
		return nil, fmt.Errorf("parse raw %s: %w", packageName, err)
	}

	return &SNILSStruct{
		numbers:  Numbers(snilsArr[:len(snilsArr)-2]),
		checkSum: checkSum(utils.SliceToInt(snilsArr[len(snilsArr)-2:])),
	}, nil
}

func NewSNILS() *SNILSStruct {
	data := &SNILSStruct{
		numbers: Numbers(utils.CodeToInts(int(utils.RandomDigits(9)))),
	}

	data.checkSum = data.calculateCheckSum()

	return data
}

func (ss *SNILSStruct) calculateCheckSum() checkSum {
	var hashSum int
	for i, v := range ss.numbers {
		hashSum += v * (len(ss.numbers) - i)
	}

	return checkSum(hashSum % 101 % 100)
}
