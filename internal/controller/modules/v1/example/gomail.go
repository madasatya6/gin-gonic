// Package v1 implements routing paths. Each services in own file.
package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/madasatya6/gin-gonic/internal/usecase/repo"
	_ "github.com/madasatya6/gin-gonic/internal/entity"
	_ "github.com/gin-contrib/sessions"
	"github.com/madasatya6/gin-gonic/pkg/gomail"
)

func (r *translationRoutes) SendEmail(c *gin.Context) {
	
	var data = make(map[string]interface{})
	data["title"] = "Learn Gin!"
	data["body"] = "Kirim email dengan Gin!"
	fileAttachment := []string{
		"app/static/img/gin.png",
	}
	
	err := gomail.SendGomail("Sendmail with Gin", "bayu.ambika@gmail.com", nil, "", "example/mail.html", data, fileAttachment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Email berhasil dikirim!",
	})
}


