package bik

import (
	"fmt"
	"strings"

	"github.com/sshaplygin/docs-code/models"
	"github.com/sshaplygin/docs-code/okato"
	"github.com/sshaplygin/docs-code/utils"
)

const packageName = "bik"

const (
	maxCountryCodeLength     = 99
	minInitConditionalNumber = 0
	maxInitConditionalNumber = 99
	minLastAccountNumbers    = 50
	maxLastAccountNumbers    = 999
)

const (
	unspecifiedCountryCode = "Неопределенный код страны"
)

var (
	// directParticipationCounty - участник платежной системы с прямым участием
	directParticipationCounty CountryCode = 0

	// indirectParticipationCounty - участник платежной системы с косвенным участием
	indirectParticipationCounty CountryCode = 1

	// notMemberClientCBRF - клиент Банка России, не являющийся участником платежной системы
	notMemberClientCBRF CountryCode = 2

	russiaCountryCode CountryCode = 4
)

var supportedCountryCodes = map[CountryCode]string{
	directParticipationCounty:   "Участник платежной системы с прямым участием",
	indirectParticipationCounty: "Участник платежной системы с косвенным участием",
	notMemberClientCBRF:         "Клиент Банка России, не являющийся участником платежной системы",
	russiaCountryCode:           "Код Российской Федерации",
}

type (
	// CountryCode Required length 2.
	CountryCode int

	// UnitConditionalNumber required length 2.
	// The conditional number of the Bank of Russia settlement network division,
	// unique within the territorial institution of the Bank of Russia,
	// in which this division of the Bank of Russia settlement network operates,
	// or the conditional number of the structural division of the Bank of Russia.
	UnitConditionalNumber int

	// LastAccountNumbers required length 3. It is last correspondent account of the bank. Possible values [050, 999]
	LastAccountNumbers int
)

const codeLength = 9

type BIKStruct struct {
	country       CountryCode
	territoryCode okato.StateCode
	unitNumber    UnitConditionalNumber
	lastNumber    LastAccountNumbers
}

func NewBIK() *BIKStruct {
	return &BIKStruct{
		country:       GenerateCountryCode(),
		territoryCode: okato.GenerateStateCode(),
		unitNumber:    GenerateUnitConditionalNumber(),
		lastNumber:    GenerateLastAccountNumbers(),
	}
}

func ParseBIK(bik string) (*BIKStruct, error) {
	if len(bik) != codeLength {
		return nil, &models.CommonError{
			Method: packageName,
			Err:    models.ErrInvalidLength,
		}
	}

	bikArr, err := utils.StrToArr(bik)
	if err != nil {
		return nil, fmt.Errorf("parse raw %s: %w", packageName, err)
	}

	return &BIKStruct{
		country:       CountryCode(utils.SliceToInt(bikArr[0:2])),
		territoryCode: okato.StateCode(utils.SliceToInt(bikArr[2:4])),
		unitNumber:    UnitConditionalNumber(utils.SliceToInt(bikArr[4:6])),
		lastNumber:    LastAccountNumbers(utils.SliceToInt(bikArr[6:])),
	}, nil
}

func (bs *BIKStruct) IsValid() (bool, error) {
	if bs == nil {
		return false, ErrNilBIK
	}

	if !bs.country.IsValid() {
		return false, ErrInvalidCountryCode
	}

	if !bs.territoryCode.IsValid() {
		return false, ErrInvalidTerritoryCode
	}

	if !bs.unitNumber.IsValid() {
		return false, ErrInvalidUnitConditionalNumber
	}

	if !bs.lastNumber.IsValid() {
		return false, ErrInvalidLastAccountNumbers
	}

	return true, nil
}

func (bs *BIKStruct) String() string {
	var res strings.Builder
	res.Grow(codeLength)

	res.WriteString(bs.country.ToString())
	res.WriteString(bs.territoryCode.ToString())
	res.WriteString(bs.unitNumber.ToString())
	res.WriteString(bs.lastNumber.ToString())

	return res.String()
}

func (bs *BIKStruct) Exists() (bool, error) {
	if bs == nil {
		return false, ErrNilBIK
	}

	_, ok := existsBIKs[bs.String()]
	return ok, nil
}

func GenerateCountryCode() CountryCode {
	// len(supportedCountryCodes)
	return russiaCountryCode
}

func GenerateUnitConditionalNumber() UnitConditionalNumber {
	return 0
}

func GenerateLastAccountNumbers() LastAccountNumbers {
	return 0
}

func (cc CountryCode) IsValid() bool {
	if cc > maxCountryCodeLength {
		return false
	}

	_, ok := supportedCountryCodes[cc]

	return ok
}

func (cc CountryCode) String() string {
	res, ok := supportedCountryCodes[cc]
	if !ok {
		return unspecifiedCountryCode
	}

	return res
}

func (cc CountryCode) ToString() string {
	return ""
}

func (ucn UnitConditionalNumber) IsValid() bool {
	return ucn >= minInitConditionalNumber && ucn <= maxInitConditionalNumber
}

func (ucn UnitConditionalNumber) ToString() string {
	return ""
}

const specialCode = 12

func (lan LastAccountNumbers) IsValid() bool {
	if lan == specialCode {
		return true
	}

	return lan >= minLastAccountNumbers && lan <= maxLastAccountNumbers
}

func (lan LastAccountNumbers) ToString() string {
	return ""
}
