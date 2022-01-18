// import (
// 	"net/http"
// 	"strings"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// )

// const expireDuration = time.Hour * 336

// // CreateToken creating token
// func CreateToken(id string) (string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["jti"] = id
// 	claims["exp"] = time.Now().Add(expireDuration).Unix()

// 	return token.SignedString([]byte(config.Configuration.SecretKey))
// }

// // GetToken get user token
// func GetToken(req *http.Request) string {
// 	cleared := strings.Replace(req.Header.Get("Authorization"), " ", "", -1)
// 	return strings.Replace(cleared, "Bearer", "", -1)
// }

// // ParseToken parse token from request
// func ParseToken(req *http.Request) (*jwt.Token, error) {
// 	tokenStr := GetToken(req)

// 	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(config.Configuration.SecretKey), nil
// 	})
// }

package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const expire = time.Hour * 336

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
