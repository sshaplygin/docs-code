package data

import "github.com/sshaplygin/docs-code/utils"

// subjectCode
const (
	subjectCodeLength = 2

	minSubjectCodeLength = 0
	maxSubjectCodeLength = 99

	// Иные территории
	defaultSubjectStrCode = "99"
	defaultSubjectCode    = 99
)

// RegionTaxServiceNumber
const (
	regionTaxServiceNumberLength = 2

	minRegionTaxServiceNumberLength  = 0
	maxRegionTaxServiceNumbereLength = 99
)

// ConstitutionSubjectCode required length 2.
// Код субъекта Российской Федерации согласно 65 статье Конституции
type ConstitutionSubjectCode int

func (csc ConstitutionSubjectCode) IsValid() bool {
	if csc < minSubjectCodeLength || csc > maxSubjectCodeLength {
		return false
	}

	_, ok := SupportedSubjectsCodes[csc]
	return ok
}

func (csc ConstitutionSubjectCode) String() string {
	_, ok := SupportedSubjectsCodes[csc]
	if !ok {
		return defaultSubjectStrCode
	}

	return utils.StrCode(int(csc), subjectCodeLength)
}

func (csc ConstitutionSubjectCode) GetName() string {
	codeName, ok := SupportedSubjectsCodes[csc]
	if !ok {
		return SupportedSubjectsCodes[defaultSubjectCode]
	}

	return codeName
}

func GenerateConstitutionSubjectCode() ConstitutionSubjectCode {
	return subjectsCodes[utils.Random(0, len(subjectsCodes)-1)]
}

type RegionTaxServiceNumber int

func (rtsn RegionTaxServiceNumber) IsValid() bool {
	if rtsn < minRegionTaxServiceNumberLength || rtsn > maxRegionTaxServiceNumbereLength {
		return false
	}

	_, ok := supportedSubjectsCodes[csc]
	return ok
}

func (rtsn RegionTaxServiceNumber) String() string {
	// todo:
	return utils.StrCode(int(rtsn), regionTaxServiceNumberLength)
}

const rootTaxServiceDepart = "0000"

type TaxRegionCode struct {
	subjectCode   ConstitutionSubjectCode
	serviceNumber RegionTaxServiceNumber
}

func (trc *TaxRegionCode) IsValid() bool {
	if trc == nil {
		return false
	}

	return trc.subjectCode.IsValid() && trc.serviceNumber.IsValid()
}

func (trc *TaxRegionCode) String() string {
	if trc == nil {
		return rootTaxServiceDepart
	}

	return trc.subjectCode.String() + trc.serviceNumber.String()
}

func (trc *TaxRegionCode) GetName() string {
	if trc == nil {
		return SupportedTaxDepartment[rootTaxServiceDepart]
	}

	codeName, ok := SupportedTaxDepartment[trc.String()]
	if !ok {
		return SupportedTaxDepartment[rootTaxServiceDepart]
	}

	return codeName
}
