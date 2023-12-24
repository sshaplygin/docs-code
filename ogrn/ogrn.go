package ogrn

import (
	"strconv"

	"github.com/sshaplygin/docs-code/models"
	"github.com/sshaplygin/docs-code/utils"
)

// Validate check to valid OGRN format
// example: input format is 1027700132195
func Validate(ogrn string) (bool, error) {
	if len(ogrn) != 13 {
		return false, &models.CommonError{
			Method: packageName,
			Err:    models.ErrInvalidLength,
		}
	}

	ogrnArr, err := utils.StrToArr(ogrn)
	if err != nil {
		return false, err
	}

	code, _ := strconv.Atoi(ogrn[:12])
	return ogrnArr[len(ogrn)-1] == code%11%10, nil
}

func Generate() string {
	panic("not implemented!")
}
