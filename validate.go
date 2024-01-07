package docs_code

import (
	"github.com/sshaplygin/docs-code/bik"
	"github.com/sshaplygin/docs-code/inn"
	"github.com/sshaplygin/docs-code/kpp"
	"github.com/sshaplygin/docs-code/ogrn"
	"github.com/sshaplygin/docs-code/ogrnip"
	"github.com/sshaplygin/docs-code/snils"
)

type ValidateFunc func(code string) (bool, error)

func Validate(docType DocType, code string) (bool, error) {
	var callFunc ValidateFunc
	switch docType {
	case BIK:
		callFunc = bik.Validate
	case INN:
		callFunc = inn.Validate
	case KPP:
		callFunc = kpp.Validate
	case OGRN:
		callFunc = ogrn.Validate
	case OGRNIP:
		callFunc = ogrnip.Validate
	case SNILS:
		callFunc = snils.Validate
	}

	if callFunc == nil {
		panic("not implemented method")
	}

	return callFunc(code)
}
