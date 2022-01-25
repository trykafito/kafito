package jwt

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Create(secretKey []byte, id string, expireAt time.Time) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["jti"] = id
	claims["exp"] = expireAt.Unix()

	return token.SignedString(secretKey)
}

func Parse(secretKey []byte, token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
}

func ParseFromRequest(secretKey []byte, req *http.Request) (*jwt.Token, error) {
	tokenStr := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")

	return Parse(secretKey, tokenStr)
}

func GetClaim(token *jwt.Token, key string) string {
	claims := token.Claims.(jwt.MapClaims)
	return claims[key].(string)
}
