package user

import (
	"strings"

	"github.com/ttacon/libphonenumber"
)

func ParsePhoneNumber(phone string, region string) (string, error) {
	number, err := libphonenumber.Parse(phone, strings.ToUpper(region))
	if err != nil {
		return "", err
	}

	result := libphonenumber.Format(number, libphonenumber.INTERNATIONAL)

	return strings.ReplaceAll(result, " ", ""), nil
}
