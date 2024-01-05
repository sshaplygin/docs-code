package kpp

import (
	"fmt"
	"strings"

	"github.com/sshaplygin/docs-code/data"
	"github.com/sshaplygin/docs-code/models"
	"github.com/sshaplygin/docs-code/utils"
)

const packageName = "kpp"

const codeLength = 9

const validateErrorTmpl = "%w: %s"

// recorderCode
const (
	recorderCodeLength = 2

	minRecorderCodeNumber = 0
	maxRecorderCodeNumber = 99
)

// reasonCode. // Could be from 0 to 9 and A to Z.
const (
	reasonCodeLength = 2

	minReasonCodeNumbers = 0
	maxReasonCodeNumbers = 99
)

// serialNumber
const (
	serialNumberLength = 3

	minSerialNumberNumbers = 0
	maxSerialNumberNumbers = 999
)

type (
	// RecorderCode required length 2.
	// код Государственной налоговой инспекции, которая осуществляла постановку на учёт организации по месту её нахождения,
	// месту нахождения её филиалов и (или)
	// представительств, расположенных на территории РФ или по месту нахождения принадлежащего ей недвижимого имущества и транспортных средств;
	RecorderCode int

	// RegistrationReason required length 2.
	// Could be from 0 to 9 and A to Z.
	// для российской организации от 01 до 50 (01 — по месту её нахождения);
	// для иностранной организации от 51 до 99;
	RegistrationReason string

	// SerialNumber required length 3.
	// порядковый номер постановки на учёт по соответствующей причине.
	SerialNumber int
)

func (rc RecorderCode) String() string {
	// TODO:
	return ""
}

func (rc RecorderCode) IsValid() bool {
	// TODO:
	return false
}

func GenerateRecorderCode() RecorderCode {
	// TODO:
	return 0
}

func (rr RegistrationReason) String() string {
	// TODO:
	return ""
}

func (rr RegistrationReason) IsValid() bool {
	// TODO:
	return false
}

func GenerateRegistrationReason() RegistrationReason {
	// TODO:
	return ""
}

func (sn SerialNumber) IsValid() bool {
	// TODO:
	return false
}

func (sn SerialNumber) String() string {
	// TODO:
	return ""
}

func GenerateSerialNumber() SerialNumber {
	// TODO:
	return 0
}

func ParseKPP(kpp string) (*KPPStruct, error) {
	if len(kpp) != codeLength {
		return nil, &models.CommonError{
			Method: packageName,
			Err:    models.ErrInvalidLength,
		}
	}

	subjectCodeArr, err := utils.StrToArr(kpp[0:2])
	if err != nil {
		return nil, fmt.Errorf("parse subject code raw %s: %w", packageName, err)
	}

	recorderCodeArr, err := utils.StrToArr(kpp[2:4])
	if err != nil {
		return nil, fmt.Errorf("parse recorder code raw %s: %w", packageName, err)
	}

	servialNumberArr, err := utils.StrToArr(kpp[6:])
	if err != nil {
		return nil, fmt.Errorf("parse serial number code raw %s: %w", packageName, err)
	}

	return &KPPStruct{
		subjectCode:  data.ConstitutionSubjectCode(utils.SliceToInt(subjectCodeArr)),
		recorderCode: RecorderCode(utils.SliceToInt(recorderCodeArr)),
		reasonCode:   RegistrationReason(kpp[4:6]),
		serialNumber: SerialNumber(utils.SliceToInt(servialNumberArr)),
	}, nil
}

type KPPStruct struct {
	subjectCode  data.ConstitutionSubjectCode
	recorderCode RecorderCode
	reasonCode   RegistrationReason
	serialNumber SerialNumber
}

func NewKPP() *KPPStruct {
	return &KPPStruct{
		subjectCode:  data.GenerateConstitutionSubjectCode(),
		recorderCode: GenerateRecorderCode(),
		reasonCode:   GenerateRegistrationReason(),
		serialNumber: GenerateSerialNumber(),
	}
}

// Примечание: для ускорения записи при голосовых переговорах используют термин «стандартный КПП» (иногда он подразумевается),
// где первые четыре цифры соответствуют ИНН, с 5 по 9 «01001».
// Организации присвоены ИНН 7750004168, КПП 7705 01001.
func (kpp *KPPStruct) IsValid() (bool, error) {
	if kpp == nil {
		return false, ErrNilKPP
	}

	if !kpp.subjectCode.IsValid() {
		return false, fmt.Errorf(validateErrorTmpl, ErrInvalidSubjectCode, kpp.subjectCode)
	}

	if !kpp.recorderCode.IsValid() {
		return false, fmt.Errorf(validateErrorTmpl, ErrInvalidRecorderCode, kpp.recorderCode)
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

	res.WriteString(kpp.subjectCode.String())
	res.WriteString(kpp.recorderCode.String())
	res.WriteString(kpp.reasonCode.String())
	res.WriteString(kpp.serialNumber.String())

	return res.String()
}
