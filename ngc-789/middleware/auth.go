package middleware

import (
	"errors"
	"fmt"
	"gin-ex/helpers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.GetHeader("Authorization")

		if authToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "authorization header required"})
			c.Abort()
			return
		}

		err := helpers.ValidateToken(authToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid authorization header"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("secretjwt")), nil
	})

	if err != nil {
		return jwt.MapClaims{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return jwt.MapClaims{}, errors.New("failed to assert map claims")
	}

}
