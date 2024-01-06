package docs_code

import (
	"github.com/sshaplygin/docs-code/bik"
	"github.com/sshaplygin/docs-code/inn"
	"github.com/sshaplygin/docs-code/kpp"
)

type GenerateFunc func() string

func Generate(docType DocType) string {
	var callFunc GenerateFunc
	switch docType {
	case BIK:
		callFunc = bik.Generate
	case INN:
		callFunc = inn.Generate
	case KPP:
		callFunc = kpp.Generate
	default:
		panic("not implemented method")
	}

	if callFunc == nil {
		panic("not implemented method")
	}

	return callFunc()
}
