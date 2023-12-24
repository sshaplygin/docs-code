package kpp

import (
	"github.com/sshaplygin/docs-code/models"
	"github.com/sshaplygin/docs-code/utils"
)

type KPP struct {
	Code         models.TaxRegionCode
	Reason       models.ReasonRegistration
	SerialNumber models.SerialNumber
}

// Validate check to valid KPP format
// example: input format is 773643301
func Validate(kpp string) (bool, error) {
	if len(kpp) != 9 {
		return false, &models.CommonError{
			Method: packageName,
			Err:    models.ErrInvalidLength,
		}
	}

	_, err := utils.StrToArr(kpp)
	if err != nil {
		return false, err
	}

	// todo: validate tax region/office models.TaxRegionCode(kpp[:4])

	_, ok := models.SupportedRegistrationReasonSet[models.RegistrationReasonCode(kpp[4:6])]
	if !ok {
		return false, ErrRegistrationReasonCode
	}

	return true, nil
}

func Generate() string {
	panic("not implemented!")
}
