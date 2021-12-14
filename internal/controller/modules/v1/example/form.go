// Package v1 implements routing paths. Each services in own file.
package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/madasatya6/gin-gonic/internal/usecase/repo"
	_ "github.com/madasatya6/gin-gonic/internal/entity"
	_ "github.com/gin-contrib/sessions"
	"github.com/utrack/gin-csrf"
	"github.com/madasatya6/gin-gonic/pkg/validation"
)

func (r *translationRoutes) GETFormMultiple(c *gin.Context) {
	//contoh penggunaan validation
	//gunakan name inputan '_csrf'
	//kalo di header gunakan 'X-CSRF-TOKEN'
	
	c.HTML(http.StatusOK, "example.validation.multiple", gin.H{
		"title": "Learning multi validation",
		"Context": c,
		"csrf_token": csrf.GetToken(c),
	})
}

func (r *translationRoutes) POSTJsonMultiple(c *gin.Context) {
	
	// Multiple form disarankan dikirimnya lewat json
	// Contoh payload data :
	//
	// {
	// 	"mahasiswa":[
	// 		{
	// 			"nama": "mada",
	// 			"alamat": "bantul",
	// 			"umur":21,
	// 			"hobbies": ["makan","main"]
	// 		},
	// 		{
	// 			"nama": "masenin",
	// 			"alamat": "jawa",
	// 			"umur":22,
	// 			"hobbies": ["makan mie","moto"]
	// 		}
	// 	]
	// }

	type Mahasiswa struct{
		Nama string `json:"nama" binding:"required"`
		Alamat string `json:"alamat" binding:"required"`
		Umur int `json:"umur" binding:"required,numeric"`
		Hobbies []string `json:"hobbies" binding:"required"`
	}
	
	type Request struct {
		Mhs []Mahasiswa `json:"mahasiswa" binding:"required"`
	}
	
	var form Request
	if err := c.ShouldBindJSON(&form); err != nil {
		list := validation.FormErrorID(c, form)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"messages": "Lengkapi data anda!",
			"data": nil,
			"errors": list,
		})
		return
	}

	// contoh validasi multi form 
	var listErrors []interface{}
	for _, mhs := range form.Mhs {
		if errs := validation.FormErrorID(c, mhs); errs != nil {
			listErrors = append(listErrors, errs)
		}
	}

	if len(listErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Mohon lengkapi data anda !",
			"data": nil,
			"errors": listErrors,
		})
		return
	} 

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"messages": "Daftar mahasiswa",
		"data": form,
		"errors": nil,
	})
	return
}

