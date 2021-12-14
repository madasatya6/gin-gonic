// Package v1 implements routing paths. Each services in own file.
package example

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/madasatya6/gin-gonic/internal/usecase/repo"
	_ "github.com/madasatya6/gin-gonic/internal/entity"
	"github.com/utrack/gin-csrf"
	"github.com/madasatya6/gin-gonic/pkg/session"
	_ "github.com/madasatya6/gin-gonic/pkg/validation"
	"github.com/gin-contrib/sessions"
)

func (r *translationRoutes) GETCaptcha(c *gin.Context) {
	//contoh penggunaan validation
	//gunakan name inputan '_csrf'
	//kalo di header gunakan 'X-CSRF-TOKEN'
	session := sessions.Default(c)
	captcha := session.Flashes("Captcha-msg")
	session.Save()

	c.HTML(http.StatusOK, "example.captcha", gin.H{
		"title": "Learning Captcha",
		"Context": c,
		"csrf_token": csrf.GetToken(c),
		"CaptchaError": captcha,
	})
}

func (r *translationRoutes) POSTCaptcha(c *gin.Context) {

	session, _ := session.Cookie(c)
	captcha := session.Values["captcha"].(string)

	log.Println(captcha, " = ", c.PostForm("captcha"))
	if captcha != c.PostForm("captcha") {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Captcha tidak cocok!",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Captcha cocok!",
		"data": captcha,
	})
}

