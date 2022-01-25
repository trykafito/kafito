package user

import (
	"errors"
	"net/http"
	"strings"

	"github.com/jeyem/passwd"
	"github.com/trykafito/kafito/pkg/jwt"
	"github.com/ttacon/libphonenumber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func LoadByRequest(secretKey []byte, r *http.Request) (*User, error) {
	t, err := jwt.ParseFromRequest(secretKey, r)
	if err != nil {
		return nil, err
	}

	id, err := primitive.ObjectIDFromHex(jwt.GetClaim(t, "jti"))
	if err != nil {
		return nil, err
	}

	return FindOne(bson.M{"_id": id})
}
