package okato

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/sshaplygin/docs-code/models"
	"github.com/sshaplygin/docs-code/utils"
)

const packageName = "okato"

const (
	stateCodeLength = 2

	minTerritoryCodeLength = 0
	maxTerritoryCodeLength = 99
)

const (
	minSecondLevelCodeLength = 0
	maxSecondLevelCodeLength = 999
)

const (
	minThirdLevelCodeLength = 0
	maxThirdLevelCodeLength = 999
)

const (
	minForuthLevelCodeLength = 0
	maxForuthLevelCodeLength = 999
)

// Структура кодового обозначения в блоке идентификации: XX XXX XXX XXX, где

// 1, 2 знаки — объекты первого уровня классификации (субъекты Российской Федерации);
// 3, 4, 5 знаки — объекты второго уровня классификации;
// 6, 7, 8 знаки — объекты третьего уровня классификации;
// КЧ — контрольное число.

const codeLength = 11

type OKATOStruct struct {
	firstLevel  StateCode
	secondLevel SecondLevelCode
	thirdLevel  ThirdLevelCode
	fourthLevel FourthLevelCode
}

func ParseOKATO(okato string) (*OKATOStruct, error) {
	okatoArr := make([]int, 0, codeLength)
	var (
		val int
		err error
	)

	for _, code := range okato {
		if code == rune(' ') {
			continue
		}

		if !unicode.IsDigit(code) {
			return nil, ErrInvalidCode
		}

		val, err = strconv.Atoi(string(code))
		if err != nil {
			return nil, fmt.Errorf("parse raw %s : %w", packageName, err)
		}

		okatoArr = append(okatoArr, val)
	}

	if len(okatoArr) != codeLength {
		return nil, &models.CommonError{
			Method: packageName,
			Err:    models.ErrInvalidLength,
		}
	}

	return &OKATOStruct{
		firstLevel:  StateCode(utils.SliceToInt(okatoArr[0:2])),
		secondLevel: SecondLevelCode(utils.SliceToInt(okatoArr[2:5])),
		thirdLevel:  ThirdLevelCode(utils.SliceToInt(okatoArr[5:8])),
		fourthLevel: FourthLevelCode(utils.SliceToInt(okatoArr[8:])),
	}, nil
}

func (ost *OKATOStruct) IsValid() (bool, error) {
	if ost == nil {
		return false, ErrNilOKATO
	}

	if !ost.firstLevel.IsValid() {
		return false, ErrFirstLevelCode
	}

	if !ost.secondLevel.IsValid() {
		return false, ErrSecondLevelCode
	}

	if !ost.thirdLevel.IsValid() {
		return false, ErrThirdLevelCode
	}

	if !ost.fourthLevel.IsValid() {
		return false, ErrFourthLevelCode
	}

	return true, nil
}

func (ost *OKATOStruct) IsExist() (bool, error) {
	panic("not implemented!")
}

// Алейского сельсовета Алейского района Алтайского края - 01 201 802 000

type StateCode int

func (s StateCode) IsValid() bool {
	return s >= minTerritoryCodeLength && s <= maxTerritoryCodeLength
}

func (s StateCode) String() string {
	_, ok := fisrtLevelCodes[s]
	if !ok {
		return StateCode(0).String()
	}

	return utils.StrCode(int(s), stateCodeLength)
}

func (s StateCode) GetName() string {
	name, ok := fisrtLevelCodes[s]
	if !ok {
		return StateCode(0).GetName()
	}

	return name
}

func GenerateStateCode() StateCode {
	return statesCodes[utils.Random(0, len(statesCodes)-1)]
}

type SecondLevelCode int

func (s SecondLevelCode) IsValid() bool {
	return s >= minSecondLevelCodeLength && s <= maxSecondLevelCodeLength
}

type ThirdLevelCode int

func (s ThirdLevelCode) IsValid() bool {
	return s >= minThirdLevelCodeLength && s <= maxThirdLevelCodeLength
}

type FourthLevelCode int

func (s FourthLevelCode) IsValid() bool {
	return s >= minForuthLevelCodeLength && s <= maxForuthLevelCodeLength
}
