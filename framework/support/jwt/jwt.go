package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/sync/singleflight"
)

type JWT struct {
	cfg Config
	sfl *singleflight.Group
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		cfg: GetConfig(),
		sfl: &singleflight.Group{},
	}
}

func (j *JWT) CreateToken(guard string, authId uint) (string, error) {
	claims := NewClaims(guard, authId, j.cfg)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.cfg.Secret)
}

func (j *JWT) ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.cfg.Secret, nil
	})
	if err == nil && token != nil {
		claims, ok := token.Claims.(*Claims)
		if ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, TokenMalformed
		}

		if ve.Errors&jwt.ValidationErrorExpired != 0 {
			// Token is expired
			return nil, TokenExpired
		}

		if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
			return nil, TokenNotValidYet
		}

		return nil, TokenInvalid
	}

	return nil, TokenInvalid
}

func (j *JWT) RefreshToken(oldToken string) (string, error) {
	key := fmt.Sprintf("JWT:REFRESH: %s", oldToken)
	res, err, _ := j.sfl.Do(key, func() (interface{}, error) {
		claims, err := j.ParseToken(oldToken)
		if err != nil {
			return "", err
		}
		return j.CreateToken(claims.Guard, claims.AuthId)
	})
	return res.(string), err
}
