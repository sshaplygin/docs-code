package bik

import (
	"fmt"

	"github.com/sshaplygin/docs-code/models"
	"github.com/sshaplygin/docs-code/utils"
)

const codeLength = 9

type BIKStruct struct {
	country       CountryCode
	territoryCode TerritoryCode
	unitNumber    UnitConditionalNumber
	lastNumber    LastAccountNumbers
}

func NewBik(bik string) (*BIKStruct, error) {
	if len(bik) != codeLength {
		return nil, &models.CommonError{
			Method: packageName,
			Err:    models.ErrInvalidLength,
		}
	}

	bikArr, err := utils.StrToArr(bik)
	if err != nil {
		return nil, fmt.Errorf("cound't parse raw bik: %w", err)
	}

	return &BIKStruct{
		country:       CountryCode(utils.SliceToInt(bikArr[0:2])),
		territoryCode: TerritoryCode(utils.SliceToInt(bikArr[2:4])),
		unitNumber:    UnitConditionalNumber(utils.SliceToInt(bikArr[4:6])),
		lastNumber:    LastAccountNumbers(utils.SliceToInt(bikArr[6:])),
	}, nil
}

func (bs *BIKStruct) IsValid() (bool, error) {
	if bs == nil {
		return false, ErrNilBik
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

// Validate check to valid BIK format.
// Example valid format is 044525225
func Validate(bik string) (bool, error) {
	bikData, err := NewBik(bik)
	if err != nil {
		return false, fmt.Errorf("create bik model: %w", err)
	}

	return bikData.IsValid()
}

func Generate() string {
	panic("not implemented!")
}
