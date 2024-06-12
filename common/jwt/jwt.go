package jwt

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	jwt.RegisteredClaims
}

func (j *JwtClaims) Valid() error {
	return nil
}

func (j *JwtClaims) TokenGenerator() {

}

func (j *JwtClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	return nil, nil
}
