package jwt

import (
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zxdstyle/liey-admin/framework/auth/authenticate"
	"golang.org/x/sync/singleflight"
	"time"
)

type JWT struct {
	cfg Config
	sfl *singleflight.Group
}

func NewJWT() *JWT {
	return &JWT{
		cfg: GetConfig(),
		sfl: &singleflight.Group{},
	}
}

func (j *JWT) CreateToken(auth authenticate.Authenticate) (string, error) {
	claims := NewClaims(auth.GuardName(), auth.GetKey(), j.cfg)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.cfg.Secret)
}

func (j *JWT) ParseToken(tokenStr string) (*Claims, error) {
	tokenStr = gstr.ReplaceByMap(tokenStr, map[string]string{
		"Bearer ": "",
		"bearer ": "",
	})
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.cfg.Secret, nil
	})
	if token == nil {
		return nil, TokenInvalid
	}

	claims, ok := token.Claims.(*Claims)
	if ok {
		return claims, j.injectError(claims, err)
	}
	return nil, TokenInvalid
}

func (j *JWT) RefreshToken(claims *Claims) (string, error) {
	key := fmt.Sprintf("JWT:REFRESH: %s", gmd5.MustEncrypt(claims))
	res, err, _ := j.sfl.Do(key, func() (interface{}, error) {
		return j.CreateToken(authenticate.NewDefaultAuthenticate(claims.AuthId, claims.Guard))
	})
	return res.(string), err
}

func (j *JWT) injectError(claims *Claims, err error) error {
	if err == nil || claims == nil {
		return nil
	}
	ve, ok := err.(*jwt.ValidationError)
	if !ok {
		return err
	}

	switch ve.Errors {
	case jwt.ValidationErrorMalformed:
		return TokenMalformed
	case jwt.ValidationErrorNotValidYet:
		return TokenNotValidYet
	case jwt.ValidationErrorExpired:
		fmt.Println(claims.RefreshAt.Unix(), claims.ExpiresAt.Unix())
		expire := uint(time.Now().Unix() - claims.ExpiresAt.Unix())
		if expire > claims.GracePeriod {
			return TokenExpired
		}
		return TokenRefresh
	default:
		return TokenInvalid
	}
}
