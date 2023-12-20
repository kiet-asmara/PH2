package helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("your_jwt_secret")

// payload struct
type JWTClaim struct {
	Store_name  string `json:"store_name"`
	Store_type  string `json:"store_type"`
	Store_email string `json:"store_email"`
	jwt.StandardClaims
}

// make jwt from payload
func GenerateJWT(store_name string, store_type string, store_email string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Store_name:  store_name,
		Store_type:  store_type,
		Store_email: store_email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtSecret)
	return
}

// take token from auth header & validate
func ValidateToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return err
	}
	return nil
}
