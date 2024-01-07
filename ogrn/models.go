package ogrn

import (
	"fmt"
	"strings"
	"time"

	"github.com/sshaplygin/docs-code/fts"
	"github.com/sshaplygin/docs-code/models"
	"github.com/sshaplygin/docs-code/utils"
)

const packageName = "ogrn"

const validateErrorTmpl = "%w: %s"

const (
	codeTypeLength = 1
)

const (
	yearsNumbersLength = 2
)

const (
	checkSumLength = 1
)

type CodeType int

func (ct CodeType) String() string {
	return utils.StrCode(int(ct), codeTypeLength)
}

func (ct CodeType) IsValid() bool {
	_, ok := supportedCodes[ct]
	return ok
}

type OGRNType uint

const (
	Physical OGRNType = iota
	Legal
	Government
)

var supportedCodes = map[CodeType]OGRNType{
	1: Legal,
	5: Legal,
	2: Government,
	4: Government,
	6: Government,
	7: Government,
	8: Government,
	9: Government,
	3: Physical,
}

type YearsNumbers int

func (yn YearsNumbers) String() string {
	return utils.StrCode(int(yn), yearsNumbersLength)
}

func (yn YearsNumbers) Ints() []int {
	res := make([]int, yearsNumbersLength)
	utils.FillSlice(utils.CodeToInts(int(yn)), res, yearsNumbersLength-1)
	return res
}

func (yn YearsNumbers) IsValid() bool {
	// 19xx: [91, 99] || 20xx: [00, now.Year%100]
	return yn >= 0 && yn <= YearsNumbers(time.Now().Year()%100) || yn >= 91 && yn <= 99
}

type serialNumbers struct {
	val int
	len int
}

func (sn serialNumbers) IsValid() bool {
	maxL := 0
	for i := 0; i < sn.len; i++ {
		maxL = maxL*10 + 9
	}
	return sn.val >= 0 && sn.val <= maxL
}

func (sn serialNumbers) String() string {
	return utils.StrCode(sn.val, sn.len)
}

func (sn serialNumbers) Ints() []int {
	res := make([]int, sn.len)
	utils.FillSlice(utils.CodeToInts(int(sn.val)), res, sn.len-1)
	return res
}

type checkSum int

func (cs checkSum) String() string {
	return utils.StrCode(int(cs), checkSumLength)
}

type OGRNStruct struct {
	code          CodeType
	yearsNumbers  YearsNumbers
	region        fts.ConstitutionRegionCode
	serialNumbers serialNumbers
	checkSum      checkSum
}

func GenerateCodeType(ogrnType OGRNType) CodeType {
	var c CodeType = 3
	if ogrnType == Legal {
		c = 1
		if utils.RandomDigits(1)%2 == 0 {
			c = 5
		}
	}
	return c
}

func GenerateYearsNumbers() YearsNumbers {
	// 19xx: [91, 99] || 20xx: [00, now.Year%100]
	min, max := 91, 99
	if utils.RandomDigits(1)%2 == 0 {
		min, max = 00, time.Now().Year()%100
	}
	return YearsNumbers(utils.Random(min, max))
}

func GenerateSerialNumbers(ogrnType OGRNType) serialNumbers {
	n := legalCodeLength
	if ogrnType == Physical {
		n = physicalCodeLength
	}

	return serialNumbers{
		val: int(utils.RandomDigits(n)),
		len: n,
	}
}

func NewOGRN(ogrnType OGRNType) *OGRNStruct {
	ogrnData := &OGRNStruct{
		code:          GenerateCodeType(ogrnType),
		yearsNumbers:  GenerateYearsNumbers(),
		region:        fts.GenerateConstitutionSubjectCode(),
		serialNumbers: GenerateSerialNumbers(ogrnType),
	}

	ogrnData.checkSum = ogrnData.calculateCheckSum()

	return ogrnData
}

const (
	legalLength    = 13
	physicalLength = 15
)

const (
	legalCodeLength    = 7
	physicalCodeLength = 9
)

func ParseOGRN(requiredType OGRNType, ogrn string) (*OGRNStruct, error) {
	if ((requiredType == Legal || requiredType == Government) && len(ogrn) != legalLength) ||
		(requiredType == Physical && len(ogrn) != physicalLength) {
		return nil, &models.CommonError{
			Method: packageName,
			Err:    models.ErrInvalidLength,
		}
	}

	ogrnArr, err := utils.StrToArr(ogrn)
	if err != nil {
		return nil, fmt.Errorf("parse raw %s: %w", packageName, err)
	}

	code := CodeType(utils.SliceToInt(ogrnArr[:1]))
	expectedCode, ok := supportedCodes[code]
	if !ok || expectedCode != requiredType {
		return nil, ErrInvalidCodeType
	}

	l := legalCodeLength
	if requiredType == Physical {
		l = physicalCodeLength
	}

	return &OGRNStruct{
		code:         code,
		yearsNumbers: YearsNumbers(utils.SliceToInt(ogrnArr[1:3])),
		region:       fts.ConstitutionRegionCode(utils.SliceToInt(ogrnArr[3:5])),
		serialNumbers: serialNumbers{
			val: utils.SliceToInt(ogrnArr[5 : len(ogrnArr)-1]),
			len: l,
		},
		checkSum: checkSum(utils.SliceToInt(ogrnArr[len(ogrnArr)-1:])),
	}, nil
}

func (o *OGRNStruct) String() string {
	codeLength := legalLength

	expectedCode := supportedCodes[o.code]
	if expectedCode == Physical {
		codeLength = physicalLength
	}

	var res strings.Builder
	res.Grow(codeLength)

	res.WriteString(o.code.String())
	res.WriteString(o.yearsNumbers.String())
	res.WriteString(o.region.String())
	res.WriteString(o.serialNumbers.String())
	res.WriteString(o.checkSum.String())

	return res.String()
}

func (o *OGRNStruct) IsValid() (bool, error) {
	if o == nil {
		return false, ErrNilOGRN
	}

	if !o.code.IsValid() {
		return false, fmt.Errorf(validateErrorTmpl, ErrInvalidCodeType, o.code)
	}

	if !o.yearsNumbers.IsValid() {
		return false, fmt.Errorf(validateErrorTmpl, ErrInvalidYearsNumbers, o.yearsNumbers)
	}

	if !o.region.IsValid() {
		return false, fmt.Errorf(validateErrorTmpl, ErrInvalidRegion, o.region)
	}

	if !o.serialNumbers.IsValid() {
		return false, fmt.Errorf(validateErrorTmpl, ErrInvalidSerialNumbers, o.serialNumbers)
	}

	return o.calculateCheckSum() == o.checkSum, nil
}

func (o *OGRNStruct) IsLegal() bool {
	t, ok := supportedCodes[o.code]
	if !ok {
		// by default OGRN for legal
		return true
	}

	return t == Legal
}

func (o *OGRNStruct) IsPhysical() bool {
	t, ok := supportedCodes[o.code]
	if !ok {
		// by default OGRN for legal
		return false
	}

	return t == Physical
}

func (o *OGRNStruct) makeSliceInts() []int {
	n := legalLength
	if o.IsPhysical() {
		n = physicalLength
	}
	res := make([]int, n-1)
	res[0] = int(o.code)

	utils.FillSlice(o.yearsNumbers.Ints(), res, 2)
	utils.FillSlice(o.region.Ints(), res, 4)
	fmt.Println("fillSlice", o.serialNumbers.Ints(), res)
	utils.FillSlice(o.serialNumbers.Ints(), res, n-2)

	return res
}

func (o *OGRNStruct) calculateCheckSum() checkSum {
	const (
		legalDelim    = 11
		physicalDelim = 13
	)

	delim := legalDelim
	if o.IsPhysical() {
		delim = physicalDelim
	}

	code := utils.SliceToInt(o.makeSliceInts())

	return checkSum(code % delim % 10)
}
