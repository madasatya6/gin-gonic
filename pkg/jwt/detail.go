package jwt 

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const BEARER_SCHEMA = "Bearer"

//get detail email from token
func Detail(c *gin.Context) jwt.MapClaims {
	
	authHeader := c.GetHeader("Authorization")

	tokenString := authHeader[len(BEARER_SCHEMA)+1:]
	if len(tokenString) == 0 {
		return nil 
	}

	token, err := JWTAuthService().ValidateToken(tokenString)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		// email: claims["name"]
		// user: claims["user"]
		// expired: claims["exp"]
		return claims
	} else {
		fmt.Printf("Error detail jwt: %w",err)
		return nil 
	}
}