// Package v1 implements routing paths. Each services in own file.
package example

import (
	"net/http"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	_ "github.com/madasatya6/gin-gonic/internal/usecase/repo"
	_ "github.com/madasatya6/gin-gonic/internal/entity"
	"github.com/utrack/gin-csrf"
	_ "github.com/madasatya6/gin-gonic/pkg/validation"
	"github.com/gin-contrib/sessions"
	"github.com/madasatya6/gin-gonic/pkg/upload"
)

type FormUpload struct{
	NamaFile string `json:"nama" form:"nama" query:"nama" binding:"required"`
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func (r *translationRoutes) GETFormUpload(c *gin.Context) {
	//contoh penggunaan validation
	//gunakan name inputan '_csrf'
	//kalo di header gunakan 'X-CSRF-TOKEN'
	session := sessions.Default(c)
	FormNama := session.Flashes("NamaFile-msg")
	FormFile := session.Flashes("File-msg")
	session.Save()

	c.HTML(http.StatusOK, "example.upload", gin.H{
		"title": "Learning upload file",
		"Context": c,
		"csrf_token": csrf.GetToken(c),
		"FormNama": FormNama,
		"FormFile": FormFile,
	})
}

func (r *translationRoutes) POSTFormUpload(c *gin.Context) {

	var request FormUpload
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": err.Error(),
		})
		return
	}

	filename, err := upload.FileValidate(c, "file", []string{"png","jpg","jpeg"}, 1, 500, true)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"message": err.Error(),
		})
		return
	}

	filename, err = upload.UploadFileAndRename(c, "file", "example", "")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Image berhasil di upload",
		"data": gin.H{
			"filename": filename,
		},
	})
}

func (r *translationRoutes) GETFormMultipleUpload(c *gin.Context) {
	//contoh penggunaan validation
	//gunakan name inputan '_csrf'
	//kalo di header gunakan 'X-CSRF-TOKEN'

	c.HTML(http.StatusOK, "example.upload.multiple", gin.H{
		"title": "Learning upload multiple file",
		"Context": c,
		"csrf_token": csrf.GetToken(c),
	})
}

func (r *translationRoutes) POSTFormMultipleUpload(c *gin.Context) {

	type FormUploadMulti struct {
		Files []*multipart.FileHeader `form:"images[]" json:"images" binding:"required"`
	}

	var request FormUploadMulti
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": err.Error(),
		})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": err.Error(),
		})
		return
	}

	files := form.File["images[]"]
	var list_error []interface{}

	for i, file := range files {
		
		//log
		r.l.Info("Indeks file upload : ", i)
		
		_, err := upload.ValidateMultipleFile(file, []string{"png","jpg","jpeg"}, 1, 500, true)
		if err != nil {
			list_error = append(list_error, err.Error())	
		}
	}

	if len(list_error) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Anda terkena validasi",
			"validated": list_error,
		})
		return 
	}

	var filesName []interface{}

	for _, file := range files {
		filename, err := upload.UploadMultipleFileAndRename(file, "example", "")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": false,
				"message": err.Error(),
			})
			return
		}
		filesName = append(filesName, filename)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Image berhasil di upload",
		"data": gin.H{
			"filename": filesName,
		},
	})
}

