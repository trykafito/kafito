package user

import (
	"errors"
	"strings"

	"github.com/jeyem/passwd"
	"github.com/ttacon/libphonenumber"
	"go.mongodb.org/mongo-driver/bson"
)

func ParsePhoneNumber(phone string, region string) (string, error) {
	number, err := libphonenumber.Parse(phone, strings.ToUpper(region))
	if err != nil {
		return "", err
	}

	result := libphonenumber.Format(number, libphonenumber.INTERNATIONAL)

	return strings.ReplaceAll(result, " ", ""), nil
}

func Auth(phone, password string) (*User, error) {
	u, err := FindOne(bson.M{"phone": phone})
	if err != nil {
		return nil, errors.New("phone or password not matched")
	}

	if !passwd.Check(password, u.Password) {
		return nil, errors.New("phone or password not matched")
	}

	return u, nil
}
