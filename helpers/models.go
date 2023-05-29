package helpers

import (
	"fmt"
	"strings"

	"github.com/biter777/countries"
)

type Model interface {
	TableName() string
}

type MobileNumber string
type Email string
type Country string

func (m *MobileNumber) Validate(countryCode string) bool {
	if strings.TrimSpace(countryCode) == "" {
		return false
	}
	return ValidateMobileNumber(string(*m), countryCode)
}

func (em *Email) Validate() error {
	return ValidateEmailDomain(string(*em))
}

func (c *Country) Validate() error {
	name := countries.ByName(string(*c))
	if strings.EqualFold(name.String(), "unknown") {
		return fmt.Errorf("%v country", name.String())
	}
	*c = Country(name.String())
	return nil
}
