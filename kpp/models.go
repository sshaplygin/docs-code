package kpp

import (
	"fmt"
	"strings"

	"github.com/sshaplygin/docs-code/fts"
	"github.com/sshaplygin/docs-code/models"
	"github.com/sshaplygin/docs-code/utils"
)

const packageName = "kpp"

const codeLength = 9

const validateErrorTmpl = "%w: %s"

// serialNumber
const (
	serialNumberLength = 3

	minSerialNumberNumbers = 0
	maxSerialNumberNumbers = 999
)

type (
	// RegistrationReason required length 2. Could be from 0 to 9 and A to Z.
	// Для российской организации от 01 до 50 (01 — по месту её нахождения).
	// Для иностранной организации от 51 до 99;
	RegistrationReason string

	// SerialNumber required length 3.
	// Порядковый номер постановки на учёт по соответствующей причине.
	SerialNumber int
)

func (rr RegistrationReason) String() string {
	return string(rr)
}

func (rr RegistrationReason) IsValid() bool {
	_, ok := supportedRegistrationsReasons[rr]
	return ok
}

func GenerateRegistrationReason() RegistrationReason {
	return registrationsReasonsCodes[utils.Random(0, len(registrationsReasonsCodes)-1)]
}

func (sn SerialNumber) IsValid() bool {
	return sn >= minSerialNumberNumbers && sn <= maxSerialNumberNumbers
}

func (sn SerialNumber) String() string {
	return utils.StrCode(int(sn), serialNumberLength)
}

func GenerateSerialNumber() SerialNumber {
	return SerialNumber(utils.RandomDigits(serialNumberLength))
}

func ParseKPP(kpp string) (*KPPStruct, error) {
	if len(kpp) != codeLength {
		return nil, &models.CommonError{
			Method: packageName,
			Err:    models.ErrInvalidLength,
		}
	}

	taxRegionCode, err := fts.ParseTaxRegionCode(kpp[0:4])
	if err != nil {
		return nil, fmt.Errorf("parse tax region code raw %s: %w", packageName, err)
	}

	servialNumberArr, err := utils.StrToArr(kpp[6:])
	if err != nil {
		return nil, fmt.Errorf("parse serial number code raw %s: %w", packageName, err)
	}

	return &KPPStruct{
		taxRegionCode: taxRegionCode,
		reasonCode:    RegistrationReason(kpp[4:6]),
		serialNumber:  SerialNumber(utils.SliceToInt(servialNumberArr)),
	}, nil
}

type KPPStruct struct {
	taxRegionCode *fts.TaxRegionCode
	reasonCode    RegistrationReason
	serialNumber  SerialNumber
}

func NewKPP() *KPPStruct {
	return &KPPStruct{
		taxRegionCode: fts.GenerateTaxRegionCode(),
		reasonCode:    GenerateRegistrationReason(),
		serialNumber:  GenerateSerialNumber(),
	}
}

func (kpp *KPPStruct) IsValid() (bool, error) {
	if kpp == nil {
		return false, ErrNilKPP
	}

	if !kpp.taxRegionCode.IsValid() {
		return false, fmt.Errorf(validateErrorTmpl, ErrInvalidTaxRegion, kpp.taxRegionCode)
	}

	if kpp.isSpecialCode() {
		return true, nil
	}

	if !kpp.reasonCode.IsValid() {
		return false, fmt.Errorf(validateErrorTmpl, ErrInvalidReasonCode, kpp.reasonCode)
	}

	if !kpp.serialNumber.IsValid() {
		return false, fmt.Errorf(validateErrorTmpl, ErrInvalidSerialNumber, kpp.serialNumber)
	}

	return true, nil
}

func (kpp *KPPStruct) String() string {
	var res strings.Builder
	res.Grow(codeLength)

	res.WriteString(kpp.taxRegionCode.String())
	res.WriteString(kpp.reasonCode.String())
	res.WriteString(kpp.serialNumber.String())

	return res.String()
}

// isSpecialCode - для ускорения записи при голосовых переговорах используют термин «стандартный КПП» (иногда он подразумевается),
// где первые четыре цифры соответствуют ИНН(не всегда), с 5 по 9 «01001».
func (kpp *KPPStruct) isSpecialCode() bool {
	const specialReasonCode RegistrationReason = "01"
	const specialSerialNumber SerialNumber = 1

	return kpp.reasonCode == specialReasonCode && kpp.serialNumber == specialSerialNumber
}
