// Package v1 implements routing paths. Each services in own file.
package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/madasatya6/gin-gonic/internal/usecase/repo"
	_ "github.com/madasatya6/gin-gonic/internal/entity"
	"github.com/gin-contrib/sessions"
	"github.com/utrack/gin-csrf"
	"github.com/madasatya6/gin-gonic/pkg/validation"
)

type FormRequest struct{
	Nama string `json:"nama" form:"nama" query:"nama" binding:"required"`
	Alamat string `json:"alamat" form:"alamat" query:"alamat" binding:"required"`
	Umur int `json:"umur" form:"umur" query:"umur" binding:"required,numeric"`
}

func (r *translationRoutes) GETValidation(c *gin.Context) {
	//contoh penggunaan validation
	//gunakan name inputan '_csrf'
	//kalo di header gunakan 'X-CSRF-TOKEN'
	session := sessions.Default(c)
	FormNama := session.Flashes("Nama-msg")
	FormAlamat := session.Flashes("Alamat-msg")
	FormUmur := session.Flashes("Umur-msg")
	session.Save()
	c.HTML(http.StatusOK, "example.validation", gin.H{
		"title": "Learning validation",
		"Context": c,
		"csrf_token": csrf.GetToken(c),
		"FormNama": FormNama,
		"FormAlamat": FormAlamat,
		"FormUmur": FormUmur,
	})
}

func (r *translationRoutes) POSTValidation(c *gin.Context) {
	var form FormRequest
	
	if err := c.ShouldBind(&form); err != nil {
		list := validation.FormErrorID(c, form)
		c.JSON(http.StatusBadRequest, gin.H{
			"messages": list,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": form,
	})
}

