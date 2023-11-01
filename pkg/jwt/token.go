package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type (
	TokenOption struct {
		AccessSecret string
		AccessExpire int64
		Fields       map[string]interface{}
	}
	Token struct {
		AccessToken  string
		AccessExpire int64
	}
)

func BuildToken(opt TokenOption) (Token, error) {
	now := time.Now().Add(-time.Minute).Unix()
	token, err := genToken(now, opt.AccessSecret, opt.AccessExpire)
	if err != nil {
		return Token{}, err
	}
	t := Token{AccessToken: token, AccessExpire: opt.AccessExpire}
	return t, nil
}

func genToken(alt int64, accessSecret string, accessExpire int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["alt"] = alt + accessExpire
	claims["exp"] = accessExpire

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(accessSecret))
	return signedString, err
}
