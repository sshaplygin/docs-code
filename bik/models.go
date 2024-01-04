package bik

import (
	"fmt"
	"strings"

	"github.com/sshaplygin/docs-code/models"
	"github.com/sshaplygin/docs-code/okato"
	"github.com/sshaplygin/docs-code/utils"
)

const packageName = "bik"

const validateErrorTmpl = "%w: %s"

const (
	countryCodeLength = 2

	minCountryCodeLength = 0
	maxCountryCodeLength = 99
)

const (
	unitConditionalNumberLength = 2

	minUnitConditionalNumber = 0
	maxUnitConditionalNumber = 99
)

const (
	lastAccountNumbersLength = 3

	minLastAccountNumbers = 50
	maxLastAccountNumbers = 999
)

const (
	unspecifiedCountryCode = "Неопределенный код страны"
)

var (
	// DirectParticipationCounty - участник платежной системы с прямым участием
	DirectParticipationCounty CountryCode = 0

	// IndirectParticipationCounty - участник платежной системы с косвенным участием
	IndirectParticipationCounty CountryCode = 1

	// NotMemberClientCBRF - клиент Банка России, не являющийся участником платежной системы
	NotMemberClientCBRF CountryCode = 2

	RussiaCountryCode CountryCode = 4
)

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

// generateOptions TODO
type generateOptions struct {
}

type GenerateOpt func(options *generateOptions)

func NewBIK(opts ...GenerateOpt) *BIKStruct {
	var options generateOptions

	for _, o := range opts {
		o(&options)
	}

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
		return false, fmt.Errorf(validateErrorTmpl, ErrInvalidCountryCode, bs.country)
	}

	if !bs.territoryCode.IsValid() {
		return false, fmt.Errorf(validateErrorTmpl, ErrInvalidTerritoryCode, bs.territoryCode)
	}

	if !bs.unitNumber.IsValid() {
		return false, fmt.Errorf(validateErrorTmpl, ErrInvalidUnitConditionalNumber, bs.unitNumber)
	}

	if !bs.lastNumber.IsValid() {
		return false, fmt.Errorf(validateErrorTmpl, ErrInvalidLastAccountNumbers, bs.lastNumber)
	}

	return true, nil
}

func (bs *BIKStruct) String() string {
	var res strings.Builder
	res.Grow(codeLength)

	res.WriteString(bs.country.String())
	res.WriteString(bs.territoryCode.String())
	res.WriteString(bs.unitNumber.String())
	res.WriteString(bs.lastNumber.String())

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
	return countryCodes[utils.Random(0, len(countryCodes)-1)]
}

func GenerateUnitConditionalNumber() UnitConditionalNumber {
	return UnitConditionalNumber(utils.Random(minUnitConditionalNumber, maxUnitConditionalNumber))
}

func GenerateLastAccountNumbers() LastAccountNumbers {
	return LastAccountNumbers(utils.Random(minLastAccountNumbers, maxLastAccountNumbers))
}

func (cc CountryCode) IsValid() bool {
	if cc < minCountryCodeLength || cc > maxCountryCodeLength {
		return false
	}

	_, ok := supportedCountryCodes[cc]
	return ok
}

func (cc CountryCode) String() string {
	_, ok := supportedCountryCodes[cc]
	if !ok {
		return RussiaCountryCode.String()
	}

	return utils.StrCode(int(cc), countryCodeLength)
}

func (cc CountryCode) GetName() string {
	codeName, ok := supportedCountryCodes[cc]
	if !ok {
		return unspecifiedCountryCode
	}

	return codeName
}

func (ucn UnitConditionalNumber) IsValid() bool {
	return ucn >= minUnitConditionalNumber && ucn <= maxUnitConditionalNumber
}

func (ucn UnitConditionalNumber) String() string {
	return utils.StrCode(int(ucn), unitConditionalNumberLength)
}

const specialCode = 12

func (lan LastAccountNumbers) IsValid() bool {
	if lan == specialCode {
		return true
	}

	return lan >= minLastAccountNumbers && lan <= maxLastAccountNumbers
}

func (lan LastAccountNumbers) String() string {
	return utils.StrCode(int(lan), lastAccountNumbersLength)
}
