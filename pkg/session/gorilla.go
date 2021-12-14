package session

import (
  	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"  
)

//session yang dipakai
var SessionStore = NewCookieStore()

func Cookie(c *gin.Context) (*sessions.Session, error) {
	create := NewCookieStore()
	sess, err := create.Get(c.Request, "id")
	return sess, err
}

//dibawah adalah configurasi session
func NewCookieStore() *sessions.CookieStore {

	authKey := []byte("my-auth-key-very-secret")
	encryptionKey := []byte("my-encryption-key-very-secret123")

	store := sessions.NewCookieStore(authKey, encryptionKey)
	store.Options.Path = "/"
	store.Options.MaxAge = 86400 * 2 //expired dalam 2 hari
	store.Options.HttpOnly = true

	return store
}

func GetFlashdata(c *gin.Context, name string) []string {
	
	session, _ := SessionStore.Get(c.Request, "fmessages")
	fm := session.Flashes(name)
	//IF we have some message

	if len(fm) > 0 {
		session.Save(c.Request, c.Writer)
		//initiate a strings slice to return messages
		var flashes []string 
		for _, fl := range fm {
			//Add message to the slice
			flashes = append(flashes, fl.(string))
		}
		
		return flashes
	}
	
	return nil
}

func SetFlashdata(c *gin.Context, name, value string){
	session, _ := SessionStore.Get(c.Request, "fmessages")
	session.AddFlash(value, name)

	session.Save(c.Request, c.Writer)
}