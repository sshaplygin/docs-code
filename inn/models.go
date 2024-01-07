package inn

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/sshaplygin/docs-code/fts"
	"github.com/sshaplygin/docs-code/models"
	"github.com/sshaplygin/docs-code/utils"
)

const packageName = "inn"

const (
	legalLength    = 10
	physicalLength = 12
)

const (
	legalSerialNumberLength    = 5
	physicalSerialNumberLength = 6
)

type INNType uint

func (t INNType) String() string {
	switch t {
	case Legal:
		return "legal"
	case ForeignLegal:
		return "foreign_legal"
	default:
		return "physical"
	}
}

const (
	Physical INNType = iota
	Legal
	// ForeignLegal start with 9909
	ForeignLegal
)

type SerialNumber struct {
	val int
	len int
}

func (sn SerialNumber) String() string {
	return utils.StrCode(sn.val, sn.len)
}

func (sn *SerialNumber) Ints() []int {
	if sn == nil {
		return nil
	}

	res := make([]int, sn.len)
	utils.FillSlice(utils.CodeToInts(sn.val), res, len(res)-1)

	return res
}

func GenerateSerailNumber(innType INNType) SerialNumber {
	if innType == Physical {
		return SerialNumber{
			val: int(utils.RandomDigits(physicalSerialNumberLength)),
			len: physicalSerialNumberLength,
		}
	}

	return SerialNumber{
		val: int(utils.RandomDigits(legalSerialNumberLength)),
		len: legalSerialNumberLength,
	}
}

type CheckSums []int

func (cs CheckSums) String() string {
	var res strings.Builder
	res.Grow(len(cs))
	for _, s := range cs {
		res.WriteString(strconv.Itoa(s))
	}
	return res.String()
}

func GenerateCheckSums(innType INNType, nums []int) CheckSums {
	checkFuncs := []checkSumFuncType{
		hash10,
	}

	shiftIdx := -1
	if innType == Physical {
		shiftIdx = -2
		checkFuncs = []checkSumFuncType{
			hash11, hash12,
		}
	}

	for _, f := range checkFuncs {
		nums = append(nums, f(nums))
	}

	fmt.Println(nums[len(nums)+shiftIdx:])
	return nums[len(nums)+shiftIdx:]
}

type INNStruct struct {
	taxRegionCode *fts.TaxRegionCode
	serialNumber  SerialNumber
	checkSums     CheckSums

	t INNType
}

func NewINN(innType INNType) *INNStruct {
	taxRegionCode := fts.GenerateTaxRegionCode()
	serialNumber := GenerateSerailNumber(innType)

	return &INNStruct{
		taxRegionCode: taxRegionCode,
		serialNumber:  serialNumber,
		checkSums:     GenerateCheckSums(innType, append(taxRegionCode.Ints(), serialNumber.Ints()...)),
	}
}

func ParseINN(inn string) (*INNStruct, error) {
	if len(inn) != legalLength && len(inn) != physicalLength {
		return nil, &models.CommonError{
			Method: packageName,
			Err:    models.ErrInvalidLength,
		}
	}

	taxRegionCode, err := fts.ParseTaxRegionCode(inn[0:4])
	if err != nil {
		return nil, fmt.Errorf("parse tax region code raw %s: %w", packageName, err)
	}

	t := Physical
	snlen := physicalSerialNumberLength
	parseIdx := len(inn) - 2
	if len(inn) == legalLength {
		snlen = legalSerialNumberLength
		t = Legal
		parseIdx = len(inn) - 1
		const foreignLegalStartWith = "9909"
		if inn[0:4] == foreignLegalStartWith {
			t = ForeignLegal
		}
	}

	serialNumberArr, err := utils.StrToArr(inn[4:parseIdx])
	if err != nil {
		return nil, fmt.Errorf("parse raw serial number %s: %w", packageName, err)
	}

	checkSums, err := getCheckSums(inn[parseIdx:])
	if err != nil {
		return nil, fmt.Errorf("get check sums value: %w", err)
	}

	return &INNStruct{
		taxRegionCode: taxRegionCode,
		serialNumber: SerialNumber{
			val: utils.SliceToInt(serialNumberArr),
			len: snlen,
		},
		checkSums: checkSums,
		t:         t,
	}, nil
}

type checkSumFuncType func(nums []int) int

func (inn *INNStruct) IsValid() (bool, error) {
	if inn == nil {
		return false, ErrNilINN
	}

	nums := append(inn.taxRegionCode.Ints(), inn.serialNumber.Ints()...)

	checkFuncs := []checkSumFuncType{
		hash10,
	}

	if inn.IsPhysical() {
		if len(nums) != physicalLength-2 {
			return false, fmt.Errorf("invalid nums length for %s type", inn.t)
		}

		if len(inn.checkSums) != 2 {
			return false, fmt.Errorf("invalid check sum length for %s type", inn.t)
		}

		checkFuncs = []checkSumFuncType{
			hash11, hash12,
		}
	}

	if inn.IsLegal() {
		if len(nums) != legalLength-1 {
			return false, fmt.Errorf("invalid nums length for %s type", inn.t)
		}

		if len(inn.checkSums) != 1 {
			return false, fmt.Errorf("invalid check sum length for %s type", inn.t)
		}
	}

	for i, f := range checkFuncs {
		if inn.checkSums[i] != f(nums) {
			return false, nil
		}
		nums = append(nums, inn.checkSums[i])
	}

	return true, nil
}

func (inn *INNStruct) String() string {
	n := physicalLength
	if inn.IsLegal() {
		n = legalLength
	}

	var res strings.Builder
	res.Grow(n)

	res.WriteString(inn.taxRegionCode.String())
	res.WriteString(inn.serialNumber.String())
	res.WriteString(inn.checkSums.String())

	return res.String()
}

func (inn *INNStruct) IsLegal() bool {
	if inn == nil {
		return false
	}

	return inn.t == Legal || inn.t == ForeignLegal
}

func (inn *INNStruct) IsPhysical() bool {
	if inn == nil {
		return false
	}

	return inn.t == Physical
}

func hash10(innArr []int) int {
	return ((2*innArr[0] + 4*innArr[1] + 10*innArr[2] + 3*innArr[3] +
		5*innArr[4] + 9*innArr[5] + 4*innArr[6] + 6*innArr[7] + 8*innArr[8]) % 11) % 10
}

func hash11(innArr []int) int {
	return ((7*innArr[0] + 2*innArr[1] + 4*innArr[2] + 10*innArr[3] + 3*innArr[4] +
		5*innArr[5] + 9*innArr[6] + 4*innArr[7] + 6*innArr[8] + 8*innArr[9]) % 11) % 10
}

func hash12(innArr []int) int {
	return ((3*innArr[0] + 7*innArr[1] + 2*innArr[2] + 4*innArr[3] +
		10*innArr[4] + 3*innArr[5] + 5*innArr[6] + 9*innArr[7] + 4*innArr[8] +
		6*innArr[9] + 8*innArr[10]) % 11) % 10
}

func getCheckSums(checkSums string) ([]int, error) {
	sums := make([]int, 0, len(checkSums))
	for _, num := range checkSums {
		if !unicode.IsDigit(num) {
			return nil, ErrInvalidCheckSumsValue
		}
		v, err := strconv.Atoi(string(num))
		if err != nil {
			return nil, fmt.Errorf("val '%c': %w", num, err)
		}
		sums = append(sums, v)
	}

	return sums, nil
}
