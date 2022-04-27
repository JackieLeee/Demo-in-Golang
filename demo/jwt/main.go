package main

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

var Logger *zap.Logger

var (
	AccessExpire int64
	AccessSecret string
)

type Claims struct {
	UID string
	jwt.RegisteredClaims
}

func BuildClaims(uid string, ttl int64) Claims {
	return Claims{
		UID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(ttl*24) * time.Hour)), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                        //签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                        //生效时间
		}}
}

func GenerateToken(userID string) (string, error) {
	claims := BuildClaims(userID, AccessExpire)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(AccessSecret))
	return tokenString, err
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(AccessSecret), nil
	}
}

func ParseToken(tokenss string) (claims *Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenss, &Claims{}, secret())

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle this token")
}

func main() {
	var err error
	Logger, err = zap.NewDevelopment()
	if err != nil {
		return
	}

	AccessSecret = "e9zw3t4hiop2mncga5vx0flky1rqudb76js8stnybwdhuvefjqyh5bv7auzq1kx7"
	AccessExpire = 7

	token, err := GenerateToken("flagship")
	if err != nil {
		Logger.Sugar().Warnf("generate token failed, err: %s", err.Error())
		return
	}
	Logger.Sugar().Debugf("token[%s]", token)
	claims, err := ParseToken(token)
	if err != nil {
		Logger.Sugar().Warnf("parse token tailed, err: %s", err.Error())
		return
	}
	Logger.Sugar().Debugf("user_id[%s]", claims.UID)
	if claims.ExpiresAt.Before(time.Now()) {
		Logger.Sugar().Warnf("this token is expired, err: %s", err.Error())
		return
	}
}
