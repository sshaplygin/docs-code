package ogrnip

import (
	"strconv"

	"github.com/sshaplygin/ru-doc-code/models"
	"github.com/sshaplygin/ru-doc-code/utils"
)

// Validate check to valid OGRNIP format
// example: input format is 304500116000157
func Validate(ogrnip string) (bool, error) {
	if len(ogrnip) != 15 {
		return false, &models.CommonError{
			Method: packageName,
			Err:    models.ErrInvalidLength,
		}
	}

	ogrnipArr, err := utils.StrToArr(ogrnip)
	if err != nil {
		return false, err
	}

	if ogrnipArr[0] != 3 && ogrnipArr[0] != 4 {
		return false, models.ErrInvalidValue
	}

	code, _ := strconv.Atoi(ogrnip[:len(ogrnip)-1])
	return ogrnipArr[len(ogrnip)-1] == code%13%10, nil
}

func Generate() string {
	panic("not implemented!")
}
