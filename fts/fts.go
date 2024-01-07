package fts

import (
	"fmt"

	"github.com/sshaplygin/docs-code/utils"
)

const packageName = "fts"

// subjectCode
const (
	subjectCodeLength = 2

	minSubjectCodeLength = 0
	maxSubjectCodeLength = 99
)

// RegionTaxServiceNumber
const (
	regionTaxServiceNumberLength = 2

	minRegionTaxServiceNumberLength  = 0
	maxRegionTaxServiceNumbereLength = 99
)

func ParseTaxRegionCode(taxRegionCode string) (*TaxRegionCode, error) {
	regionCodeArr, err := utils.StrToArr(taxRegionCode[:2])
	if err != nil {
		return nil, fmt.Errorf("parse subject code raw %s: %w", packageName, err)
	}

	serviceNumberArr, err := utils.StrToArr(taxRegionCode[2:])
	if err != nil {
		return nil, fmt.Errorf("parse subject code raw %s: %w", packageName, err)
	}

	return &TaxRegionCode{
		subjectCode:   ConstitutionRegionCode(utils.SliceToInt(regionCodeArr)),
		serviceNumber: RegionTaxServiceNumber(utils.SliceToInt(serviceNumberArr)),
	}, nil
}

type TaxRegionCode struct {
	subjectCode   ConstitutionRegionCode
	serviceNumber RegionTaxServiceNumber
}

func (trc *TaxRegionCode) Ints() []int {
	if trc == nil {
		return nil
	}

	res := make([]int, subjectCodeLength+regionTaxServiceNumberLength)

	utils.FillSlice(utils.CodeToInts(int(trc.subjectCode)), res, subjectCodeLength-1)
	utils.FillSlice(utils.CodeToInts(int(trc.serviceNumber)), res, len(res)-1)

	return res
}

func (trc *TaxRegionCode) IsValid() bool {
	if trc == nil {
		return false
	}

	return trc.subjectCode.IsValid() && trc.serviceNumber.IsValid(trc.subjectCode)
}

func (trc *TaxRegionCode) String() string {
	const rootTaxServiceDepart = "0000"

	if trc == nil {
		return rootTaxServiceDepart
	}

	return trc.subjectCode.String() + trc.serviceNumber.String()
}

func (trc *TaxRegionCode) GetName() string {
	const rootTaxServiceDepart = 0

	if trc == nil {
		return SupportedTaxDepartments[rootTaxServiceDepart].Name
	}

	region, ok := SupportedTaxDepartments[trc.subjectCode]
	if !ok {
		return SupportedTaxDepartments[rootTaxServiceDepart].Name
	}

	depart, ok := region.Branches[trc.serviceNumber]
	if !ok {
		return region.Name
	}

	return depart
}

func GenerateTaxRegionCode() *TaxRegionCode {
	subjectCode := GenerateConstitutionSubjectCode()

	return &TaxRegionCode{
		subjectCode:   subjectCode,
		serviceNumber: GenerateRegionTaxServiceNumber(subjectCode),
	}
}

// ConstitutionRegionCode required length 2.
// Код субъекта Российской Федерации согласно 65 статье Конституции
type ConstitutionRegionCode int

func (csc ConstitutionRegionCode) IsValid() bool {
	if csc < minSubjectCodeLength || csc > maxSubjectCodeLength {
		return false
	}

	_, ok := SupportedRegionsCodes[csc]
	return ok
}

func (csc ConstitutionRegionCode) String() string {
	// Иные территории
	const defaultSubjectStrCode = "99"

	_, ok := SupportedRegionsCodes[csc]
	if !ok {
		return defaultSubjectStrCode
	}

	return utils.StrCode(int(csc), subjectCodeLength)
}

func (csc ConstitutionRegionCode) GetName() string {
	// Иные территории
	const defaultSubjectCode = 99

	codeName, ok := SupportedRegionsCodes[csc]
	if !ok {
		return SupportedRegionsCodes[defaultSubjectCode]
	}

	return codeName
}

func (csc ConstitutionRegionCode) Ints() []int {
	res := make([]int, subjectCodeLength)
	utils.FillSlice(utils.CodeToInts(int(csc)), res, len(res)-1)
	return res
}

func GenerateConstitutionSubjectCode() ConstitutionRegionCode {
	return supportedRegionsCodes[utils.Random(0, len(supportedRegionsCodes)-1)]
}

type RegionTaxServiceNumber int

func (rtsn RegionTaxServiceNumber) IsValid(sc ConstitutionRegionCode) bool {
	if rtsn < minRegionTaxServiceNumberLength || rtsn > maxRegionTaxServiceNumbereLength {
		return false
	}

	_, ok := SupportedTaxDepartments[sc].Branches[rtsn]
	return ok
}

func (rtsn RegionTaxServiceNumber) String() string {
	return utils.StrCode(int(rtsn), regionTaxServiceNumberLength)
}

func GenerateRegionTaxServiceNumber(sc ConstitutionRegionCode) RegionTaxServiceNumber {
	departs, ok := SupportedTaxDepartments[sc]
	if !ok || len(departs.Branches) == 0 {
		return 0
	}

	departsCodes := make([]RegionTaxServiceNumber, 0, len(departs.Branches))
	for code := range departs.Branches {
		departsCodes = append(departsCodes, code)
	}

	return departsCodes[utils.Random(0, len(departsCodes)-1)]
}
