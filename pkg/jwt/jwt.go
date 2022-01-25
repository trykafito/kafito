package jwt

import (
	"fmt"
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
	fmt.Println(req.Header.Get("Authorization"))
	tokenStr := strings.Replace(req.Header.Get("Authorization"), " ", "", -1)
	tokenStr = strings.Replace(tokenStr, "Bearer", "", -1)
	fmt.Println(tokenStr)

	return Parse(secretKey, tokenStr)
}

func GetClaim(token *jwt.Token, key string) string {
	claims := token.Claims.(jwt.MapClaims)
	return claims[key].(string)
}
