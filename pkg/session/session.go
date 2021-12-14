package session

import (
	"github.com/gin-contrib/sessions"
  	"github.com/gin-contrib/sessions/cookie"
  	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge:   60 * 60 * 24}) // expire in a day
	r.Use(sessions.Sessions("mysession", store))
}
