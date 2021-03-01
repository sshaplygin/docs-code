package kpp

import (
	ru_doc_code "github.com/mrfoe7/ru-doc-code"
)

type KPP struct {
	Code         ru_doc_code.TaxRegionCode
	Reason       ru_doc_code.ReasonRegistration
	SerialNumber ru_doc_code.SerialNumber
}

// Validate check to valid KPP format
// example: input format is 773643301
func Validate(kpp string) (bool, error) {
	if len(kpp) != 9 {
		return false, ru_doc_code.ErrInvalidKPPLength
	}

	_, err := ru_doc_code.StrToArr(kpp)
	if err != nil {
		return false, err
	}

	// todo: validate tax region/office ru_doc_code.TaxRegionCode(kpp[:4])

	_, ok := ru_doc_code.SupportedRegistrationReasonSet[ru_doc_code.RegistrationReasonCode(kpp[4:6])]
	if !ok {
		return false, ru_doc_code.ErrInvalidRegistrationReasonCode
	}

	return true, nil
}
