package inn

import "github.com/sshaplygin/docs-code/models"

const packageName = "packageName"

const (
	lengthLegal    = 10
	lengthPhysical = 12
)

type SerialNumber int

type INNStruct struct {
	taxCode      models.TaxRegionCode
	serialNumber SerialNumber
	hash10       uint
	hash11       uint
	hash12       uint
}
