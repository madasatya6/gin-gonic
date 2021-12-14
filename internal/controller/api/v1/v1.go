// Package v1 implements routing paths. Each services in own file.
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/madasatya6/gin-gonic/internal/usecase/repo"
	"github.com/madasatya6/gin-gonic/internal/usecase"
	"github.com/madasatya6/gin-gonic/pkg/logger"
	"github.com/madasatya6/gin-gonic/pkg/jwt"
	"github.com/madasatya6/gin-gonic/internal/entity"
	_ "github.com/gin-contrib/sessions"
)

type translationRoutes struct{
	t usecase.Translation
	l logger.Interface
}

func Init(handler *gin.RouterGroup, t usecase.Translation, l logger.Interface){
	route := &translationRoutes{t,l}
	/* use csrf
	* In order to make use of gin-csrf you'll need to either:

    - Pass the token using a form descriptor _csrf
    - Add the parameter to your request _csrf + token
    - Include either keys X-CSRF-TOKEN or X-XSRF-TOKEN in your header with the value being the token
	*/
	v1 := handler.Group("/v1")
	{
		v1.POST("/login", route.JwtLogin)
		auth := v1.Group("/auth")
		{
			auth.Use(jwt.AuthorizeJWT())
			auth.GET("/detail", route.JwtDetail)
		}
	}
}

func (r *translationRoutes) JwtLogin(c *gin.Context) {
	
	var request entity.LoginCredentials
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": err.Error(),
		})
		return 
	}

	if request.Email != "admin@gmail.com" || request.Password != "123456" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Anda belum terdaftar!",
		})
		return 
	}

	service := jwt.JWTAuthService()
	token := service.GenerateToken(request.Email, true)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Login JWT berhasil!",
		"data": gin.H{
			"token": token,
		},
	})
}

func (r *translationRoutes) JwtDetail(c *gin.Context) {

	claims := jwt.Detail(c)
	if claims == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Token anda salah!",
			"data": nil,
		})
		return 
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Detail akun anda",
		"data": claims,
	})
}
