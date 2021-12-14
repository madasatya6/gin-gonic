package captcha

import (
	"image/color"
	"image/png"

	"github.com/gin-gonic/gin"
	"github.com/afocus/captcha"
	"github.com/madasatya6/gin-gonic/pkg/session"
	_ "github.com/gin-contrib/sessions"
)

//simple captcha
func CreateCaptcha() *captcha.Captcha {
	cap := captcha.New()

	if err := cap.SetFont("app/static/fonts/comic.ttf"); err != nil {
		panic(err.Error())
	}

	cap.SetSize(128, 64)
	cap.SetDisturbance(captcha.MEDIUM)
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	return cap 
}

//route captcha gin gonic
func GenerateCaptcha(c *gin.Context) {
	
	var cap *captcha.Captcha = CreateCaptcha()
	img, str := cap.Create(6, captcha.ALL)
	
	println("akses captcha:", str)
	//save in session
	session, _ := session.Cookie(c)
	session.Values["captcha"] = str
	session.Save(c.Request, c.Writer)
	
	png.Encode(c.Writer, img)
}

