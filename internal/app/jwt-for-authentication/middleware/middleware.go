package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"net/http"
)

const (
	AccessToken = "access-token"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, _ := c.Cookie(AccessToken)
		if c.GetHeader(AccessToken) != "" {
			tokenString = c.GetHeader(AccessToken)
		}

		if tokenString != "" {
			verifyBytes, err := ioutil.ReadFile("/Users/lap01651/nogopy/jwt-for-authentication/jwtRS256.key.pub")
			// TODO: Please set private key in your env, not in your code!
			verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
			if err != nil {
				abort(c)
				return
			}

			token, err := new(jwt.Parser).ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
				return verifyKey, nil
			})
			if err != nil {
				abort(c)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				c.Set("user_id", claims["user_id"])
				c.Next()
				return
			}
		}

		abort(c)
	}
}

func abort(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": gin.H{
			"message": "Unauthorized!",
			"code":    http.StatusUnauthorized,
		},
	})
}
