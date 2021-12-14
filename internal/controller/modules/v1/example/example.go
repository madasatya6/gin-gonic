// Package v1 implements routing paths. Each services in own file.
package example

import (
	"net/http"
	"context"

	"github.com/gin-gonic/gin"
	_ "github.com/madasatya6/gin-gonic/internal/usecase/repo"
	"github.com/madasatya6/gin-gonic/internal/usecase"
	"github.com/madasatya6/gin-gonic/pkg/logger"
	csrfc "github.com/madasatya6/gin-gonic/pkg/csrf"
	_ "github.com/madasatya6/gin-gonic/internal/entity"
	"github.com/gin-contrib/sessions"
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
	handler.POST("/json-multiple", route.POSTJsonMultiple)
	handler.Use(csrfc.New())
	
	handler.GET("/crud-read", route.CRUDRead)
	handler.GET("/session-get", route.GetSession)
	handler.GET("/session-set", route.SetSession)
	handler.GET("/session-flash-set", route.SetFlashSession)
	handler.GET("/session-flash-get", route.GetFlashSession)
	handler.GET("/view", route.View)
	handler.GET("/form", route.GETValidation)
	handler.POST("/validate", route.POSTValidation)
	handler.GET("/language", route.GETLanguage)
	handler.GET("/pagination", route.GETPagination)
	handler.GET("/pagination-view", route.GETPaginationView)
	handler.GET("/mail", route.SendEmail)
	handler.GET("/google-translate", route.GoogleTranslate)
	handler.GET("/upload", route.GETFormUpload)
	handler.POST("/upload-post", route.POSTFormUpload)
	handler.GET("/upload-multiple", route.GETFormMultipleUpload)
	handler.POST("/upload-multiple-post", route.POSTFormMultipleUpload)
	handler.GET("/captcha", route.GETCaptcha)
	handler.POST("/captcha-post", route.POSTCaptcha)
	handler.GET("/form-multiple", route.GETFormMultiple)
}

func (r *translationRoutes) CRUDRead(c *gin.Context) {
	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repo :=  r.t.Setting()

	data, err := repo.All(ctx)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Data example",
		"data": data,
	})
}

func (r *translationRoutes) SetSession(c *gin.Context) {
	//contoh set session
	var email = c.Query("email")
	session := sessions.Default(c)
	session.Set("email", email)
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"email": email,
	})
}

func (r *translationRoutes) GetSession(c *gin.Context) {
	//contoh get session
	session := sessions.Default(c)
	email := session.Get("email")

	c.PureJSON(http.StatusOK, gin.H{
		"email": email,
	}) 
	return
}

func (r *translationRoutes) SetFlashSession(c *gin.Context) {
	//contoh penggunaan flash session
	session := sessions.Default(c)
	session.AddFlash("Data berhasil disimpan!", "Notifikasi")
	session.Save()
	
	c.Redirect(http.StatusFound, "/example/session-flash-get")
} 

func (r *translationRoutes) GetFlashSession(c *gin.Context) {
	//contoh penggunaan get flash session
	session := sessions.Default(c)
	notif := session.Flashes("Notifikasi")
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"Notifikasi": notif,
	})
}

func (r *translationRoutes) View(c *gin.Context) {
	//contoh penggunaan view
	c.HTML(http.StatusOK, "example", gin.H{
		"title": "Belajar template",
		"body": "Content Example",
		"footer": "This is Footer",
	})
}

