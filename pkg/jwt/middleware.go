package jwt 

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//route.Use(AuthorizeJWT())
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")

		if strings.Index(authHeader, BEARER_SCHEMA) == -1 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": false,
				"message": "Format token anda salah!",
				"data": nil,
			})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA)+1:]

		token, err := JWTAuthService().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": false,
				"message": "Token anda tidak terdaftar!",
				"data": nil,
			})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}