package fts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseTaxRegionCode(t *testing.T) {
	trc, err := ParseTaxRegionCode("7746")
	require.NoError(t, err)

	assert.Equal(t, "7746", trc.String())
	assert.Equal(t, []int{7, 7, 4, 6}, trc.Ints())
}

func TestParseTaxRegionCode_Invalid(t *testing.T) {
	_, err := ParseTaxRegionCode("77XX")
	require.Error(t, err)
}

func TestTaxRegionCode_Nil(t *testing.T) {
	var trc *TaxRegionCode

	assert.Nil(t, trc.Ints())
	assert.False(t, trc.IsValid())
	assert.Equal(t, "0000", trc.String())
	assert.Equal(t, SupportedTaxDepartments[0].Name, trc.GetName())
}

func TestConstitutionRegionCode(t *testing.T) {
	const adygea ConstitutionRegionCode = 1

	assert.True(t, adygea.IsValid())
	assert.Equal(t, "01", adygea.String())
	assert.Equal(t, SupportedRegionsCodes[adygea], adygea.GetName())
	assert.Equal(t, []int{0, 1}, adygea.Ints())
}

func TestConstitutionRegionCode_Invalid(t *testing.T) {
	const unknown ConstitutionRegionCode = 100

	assert.False(t, unknown.IsValid())
	// Unknown/out-of-range subjects fall back to the "other territories" code.
	assert.Equal(t, "99", unknown.String())
}

func TestRegionTaxServiceNumber_String(t *testing.T) {
	assert.Equal(t, "05", RegionTaxServiceNumber(5).String())
	assert.Equal(t, "42", RegionTaxServiceNumber(42).String())
}

func TestGenerateConstitutionSubjectCode_AlwaysValid(t *testing.T) {
	for range 100 {
		require.True(t, GenerateConstitutionSubjectCode().IsValid())
	}
}

func TestGenerateTaxRegionCode_RoundTrip(t *testing.T) {
	for range 100 {
		trc := GenerateTaxRegionCode()
		require.NotNil(t, trc)
		require.True(t, trc.IsValid())

		parsed, err := ParseTaxRegionCode(trc.String())
		require.NoError(t, err)
		assert.Equal(t, trc.Ints(), parsed.Ints())
	}
}
