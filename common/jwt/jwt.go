package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var signingKey = []byte("backend.com53cr3t")

type JwtClaims struct {
	jwt.RegisteredClaims
}

func (j *JwtClaims) Valid(token string) (*JwtClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if !jwtToken.Valid {
		//Logger.Warn("valid token err")
		return nil, fmt.Errorf("invalid token")
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		//Logger.Warn("That's not even a token")
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		//Logger.Warn("Invalid signature")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		//Logger.Warn("Timing is everything")
		return nil, fmt.Errorf("登录失效，请重新登录")
	} else {
		//log.Println("Couldn't handle this token:", err)
	}
	if err != nil {
		return nil, err
	}
	claims, ok := jwtToken.Claims.(*JwtClaims)
	if !ok {
		return nil, err
	}
	return claims, nil
}

// TokenGenerator 生成token
func (j *JwtClaims) TokenGenerator(userID string) (expireTime time.Time, t string, err error) {
	expireTime = time.Now().Add(24 * time.Hour)
	j.ExpiresAt = jwt.NewNumericDate(expireTime)
	j.IssuedAt = jwt.NewNumericDate(time.Now())
	j.NotBefore = jwt.NewNumericDate(time.Now())
	j.Issuer = "cs.com"
	j.ID = uuid.NewString()
	j.Subject = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j)
	signedString, err := token.SignedString(signingKey)
	if err != nil {
		return time.Time{}, "", err
	}
	return expireTime, signedString, nil
}

// RefreshTokenGenerator 生成刷新token
func (j *JwtClaims) RefreshTokenGenerator(userID string) (string, error) {
	expireTime := time.Now().Add(24 * time.Hour * 30)
	j.ExpiresAt = jwt.NewNumericDate(expireTime)
	j.IssuedAt = jwt.NewNumericDate(time.Now())
	j.NotBefore = jwt.NewNumericDate(time.Now())
	j.Issuer = "cs.com"
	j.ID = uuid.NewString()
	j.Subject = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j)
	signedString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

// RefreshToken 调用接口使用refreshToken 换取 新的accessToken
func (j *JwtClaims) RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	// refresh token无效直接返回
	if _, err = jwt.Parse(rToken, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	}); err != nil {
		return
	}

	// 从旧access token中解析出claims数据	解析出payload负载信息
	var claims JwtClaims
	_, err = jwt.ParseWithClaims(aToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	// 当access token是过期错误 并且 refresh token没有过期时就创建一个新的access token
	if errors.Is(err, jwt.ErrTokenExpired) {
		refreshToken, err := j.RefreshTokenGenerator(claims.Subject)
		_, accessToken, err := j.TokenGenerator(claims.Subject)
		return accessToken, refreshToken, err
	}
	return
}

func (j *JwtClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	return j.IssuedAt, nil
}
