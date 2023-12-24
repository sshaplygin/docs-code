package bik

const packageName = "bik"

const (
	maxCountryCodeLength     = 100
	maxTerritoryCodeLength   = 100
	maxInitConditionalNumber = 100
	minLastAccountNumbers    = 50
	maxLastAccountNumbers    = 1000
)

var (
	// directParticipationCounty - участник платежной системы с прямым участием
	directParticipationCounty CountryCode = 0

	// indirectParticipationCounty - участник платежной системы с косвенным участием
	indirectParticipationCounty CountryCode = 1

	// notMemberClientCBRF - клиент Банка России, не являющийся участником платежной системы
	notMemberClientCBRF CountryCode = 2

	russiaCountryCode CountryCode = 4
)

var supportedCountryCodes = map[CountryCode]string{
	directParticipationCounty:   "Участник платежной системы с прямым участием",
	indirectParticipationCounty: "Участник платежной системы с косвенным участием",
	notMemberClientCBRF:         "Клиент Банка России, не являющийся участником платежной системы",
	russiaCountryCode:           "Код Российской Федерации",
}

type (
	// CountryCode Required length 2.
	CountryCode int

	// TerritoryCode OKATO code. Required length 2.
	TerritoryCode int

	// UnitConditionalNumber required length 2.
	// The conditional number of the Bank of Russia settlement network division,
	// unique within the territorial institution of the Bank of Russia,
	// in which this division of the Bank of Russia settlement network operates,
	// or the conditional number of the structural division of the Bank of Russia.
	UnitConditionalNumber int

	// LastAccountNumbers required length 3. It is last correspondent account of the bank. Possible values 050 до 999
	LastAccountNumbers int
)

func (cc CountryCode) IsValid() bool {
	if cc > maxCountryCodeLength {
		return false
	}

	_, ok := supportedCountryCodes[cc]

	return ok
}

func (tc TerritoryCode) IsValid() bool {
	return tc < maxTerritoryCodeLength
}

func (ucn UnitConditionalNumber) IsValid() bool {
	return ucn < maxInitConditionalNumber
}

const specialCode = 12

func (lan LastAccountNumbers) IsValid() bool {
	if lan == specialCode {
		return true
	}

	return lan >= minLastAccountNumbers && lan < maxLastAccountNumbers
}
