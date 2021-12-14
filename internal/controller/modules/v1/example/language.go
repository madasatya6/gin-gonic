// Package v1 implements routing paths. Each services in own file.
package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/madasatya6/gin-gonic/internal/usecase/repo"
	_ "github.com/madasatya6/gin-gonic/internal/entity"
	"github.com/utrack/gin-csrf"
	"github.com/gin-contrib/sessions"
	local "github.com/madasatya6/gin-gonic/pkg/language"
)

func (r *translationRoutes) GETLanguage(c *gin.Context) {
	//contoh penggunaan validation
	//gunakan name inputan '_csrf'
	//kalo di header gunakan 'X-CSRF-TOKEN'
	var language string
	session := sessions.Default(c)
	bahasa := c.Query("language")
	if session.Get("language") == nil {
		language = ""
	} else {
		language = session.Get("language").(string)
	}
	local.Language = &language

	if bahasa != "" {
		session.Set("language", bahasa)
		session.Save()
	}
	c.HTML(http.StatusOK, "example.language", gin.H{
		"title": "Learning Local Translation",
		"Context": c,
		"csrf_token": csrf.GetToken(c),
		"language": language,
	})
}


