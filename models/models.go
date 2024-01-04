package models

import "github.com/sshaplygin/docs-code/utils"

// subjectCode
const (
	subjectCodeLength = 2

	minSubjectCodeLength = 0
	maxSubjectCodeLength = 99
)

// ConstitutionSubjectCode required length 2.
// Код субъекта Российской Федерации согласно 65 статье Конституции
type ConstitutionSubjectCode int

func (csc ConstitutionSubjectCode) IsValid() bool {
	if csc < minSubjectCodeLength || csc > maxSubjectCodeLength {
		return false
	}

	_, ok := supportedSubjectsCodes[csc]
	return ok
}

func (csc ConstitutionSubjectCode) String() string {
	_, ok := supportedCountryCodes[cc]
	if !ok {
		return RussiaCountryCode.String()
	}

	return utils.StrCode(int(cc), countryCodeLength)
}

func (csc ConstitutionSubjectCode) GetName() string {
	codeName, ok := supportedCountryCodes[cc]
	if !ok {
		return unspecifiedCountryCode
	}

	return codeName
}

type RegionTaxServiceNumber int

func (rtn RegionTaxServiceNumber) IsValid() bool {
	if csc < minSubjectCodeLength || csc > maxSubjectCodeLength {
		return false
	}

	_, ok := supportedSubjectsCodes[csc]
	return ok
}

func (rtn RegionTaxServiceNumber) String() string {
	_, ok := supportedCountryCodes[cc]
	if !ok {
		return RussiaCountryCode.String()
	}

	return utils.StrCode(int(cc), countryCodeLength)
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
