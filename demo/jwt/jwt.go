package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	// AccessExpire 过期时间
	AccessExpire int64
	// AccessSecret 密钥
	AccessSecret string
)

type Claims struct {
	// 用户id
	UID string
	// 注册Claims
	jwt.RegisteredClaims
}

// BuildClaims 创建Claims
func BuildClaims(uid string, ttl int64) Claims {
	return Claims{
		UID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(ttl*24) * time.Hour)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                        // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                        // 生效时间
		}}
}

// GenerateToken 生成token
func GenerateToken(userID string) (string, error) {
	claims := BuildClaims(userID, AccessExpire)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(AccessSecret))
	return tokenString, err
}

// secret 获取secret
func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(AccessSecret), nil
	}
}

// ParseToken 解析token
func ParseToken(tokenss string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenss, &Claims{}, secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			switch {
			case ve.Errors&jwt.ValidationErrorMalformed != 0:
				return nil, errors.New("that's not even a token")
			case ve.Errors&jwt.ValidationErrorExpired != 0:
				return nil, errors.New("token is expired")
			case ve.Errors&jwt.ValidationErrorNotValidYet != 0:
				return nil, errors.New("token not active yet")
			default:
				return nil, errors.New("couldn't handle this token")
			}
		}
		return nil, fmt.Errorf("unknown error, err: %s", err.Error())
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle this token")
}
